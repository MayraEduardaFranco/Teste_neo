/*package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)
func main ()
func OpenConn() (*sql.db,error){
	db, err:= sql.Open("postgres", "host=localhost port =5432 user=postgres password=Mepf@2002 dbname=postgres sslmode=disable")
	if err != nil {
		panic (err)
	}
	err = db.Ping ()
	return db, err
}
type Categoria struct {
    ID   int64
    Nome string
}
func Insert(c Categoria) (id int64, err error) {
    conn, err := db.OpenConn()
    if err != nil {
        return
    }
    defer conn.Close()

    sql := `INSERT INTO categorias (nome) VALUES ($1) RETURNING id`

    err = conn.QueryRow(sql, c.Nome).Scan(&id)

    return id
}
// categoria especifica
func Get(id int64) (c Categoria, err error) {
    conn, err := db.OpenConn()
    if err != nil {
        return
    }
    defer conn.Close()

    row := conn.QueryRow(`SELECT * FROM categorias WHERE id=$1`, id)

    err = row.Scan(&c.ID, &c.Nome)

    return
}

// todas as categorias
func GetAll() (sc []Categoria, err error) {
    conn, err := db.OpenConn()
    if err != nil {
        return
    }
    defer conn.Close()

    rows, err := db.Query(`SELECT * FROM categorias`)
    if err != nil {
        return
    }
    defer rows.Close()

    for rows.Next() {
        var c Categoria

        err = row.Scan(&c.ID, &c.Nome)
        if err != nil {
            continue
        }

        sc = append(sc, c)
    }

    return
}
func Update(id int64, c Categoria) (int64, error) {
    conn, err := db.OpenConn()
    if err != nil {
        return
    }
    defer conn.Close()

    res, err := db.Exec(`UPDATE categorias SET nome=$2 WHERE id=$1`, id, c.Nome)
    if err != nil {
        return err
    }

    return res.RowsAffected()
}
func Delete(id int64) (int64, error) {
    conn, err := db.OpenConn()
    if err != nil {
        return
    }
    defer conn.Close()

    res, err := db.Exec(`DELETE FROM users WHERE id=$1`, id)
    if err != nil {
        return err
    }

    return res.RowsAffected()
}

package main
import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	_"github.com/lib/pq"
)
func main (){
	//conectar banco de dados Postgre
	connStr := "user=postgre dbname=postgre password=Mepf@2002 sslmode=disable"
	db, err := sql.Open ("postgres",connStr)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	//criar a tabela
	createTableQuery := 
     sua_tabela (
        id SERIAL PRIMARY KEY,
        coluna1 TEXT,
        coluna2 TEXT,
        coluna3 TEXT
    );
	_, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }

    // Abrir o arquivo CSV
    file, err := os.Open("arquivo.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Ler os dados do CSV
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Inserir os dados na tabela
    insertQuery := INSERT INTO sua_tabela (coluna1, coluna2, coluna3) VALUES ($1, $2, $3)
    for _, record := range records {
        _, err := db.Exec(insertQuery, record[0], record[1], record[2])
        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("Dados inseridos com sucesso!")
}
package main

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "reflect"
    _ "github.com/lib/pq"
    "strings"
)

// Conexão com o PostgreSQL
const connStr = "user=postgres dbname=postgres password=Mepf@2002 sslmode=disable"

// Função para definir tipos PostgreSQL com base nos tipos do Go
func getPostgresType(goType string) string {
    switch goType {
    case "string":
        return "TEXT"
    case "int", "int64":
        return "INTEGER"
    case "float64":
        return "FLOAT"
    default:
        return "TEXT"
    }
}

// Função para criar uma tabela com base nos dados do CSV
func createTable(db *sql.DB, tableName string, headers []string, sampleRow []string) error {
    var columns []string

    // Determinar tipos de dados com base na primeira linha de dados (sampleRow)
    for i, header := range headers {
        goType := reflect.TypeOf(sampleRow[i]).Name()
        pgType := getPostgresType(goType)
        columns = append(columns, fmt.Sprintf("%s %s", header, pgType))
    }

    createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columns, ", "))
    _, err := db.Exec(createTableSQL)
    return err
}

// Função para converter []string em []interface{}
func stringSliceToInterface(slice []string) []interface{} {
    iSlice := make([]interface{}, len(slice))
    for i, v := range slice {
        iSlice[i] = v
    }
    return iSlice
}

// Função para ler dados do CSV e inserir no PostgreSQL
func importCSVToPostgres(db *sql.DB, filePath string, tableName string) error {
    // Abrir arquivo CSV
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // Ler CSV
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    // Pegar cabeçalhos e primeira linha para criar a tabela
    headers := records[0]
    if len(records) < 2 {
        return fmt.Errorf("CSV está vazio")
    }
    sampleRow := records[1]

    // Criar a tabela no PostgreSQL
    err = createTable(db, tableName, headers, sampleRow)
    if err != nil {
        return err
    }

    // Inserir os dados no PostgreSQL
    insertSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(headers, ", "), placeholders(len(headers)))
    for _, row := range records[1:] {
        // Converter []string para []interface{}
        _, err := db.Exec(insertSQL, stringSliceToInterface(row)...)
        if err != nil {
            return err
        }
    }
    return nil
}

// Função para gerar placeholders de SQL ($1, $2, etc.)
func placeholders(n int) string {
    var ph []string
    for i := 1; i <= n; i++ {
        ph = append(ph, fmt.Sprintf("$%d", i))
    }
    return strings.Join(ph, ", ")
}

func main() {
    // Conectar ao banco de dados PostgreSQL
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Definir os arquivos CSV e os nomes das tabelas
    csvFiles := map[string]string{
        "C:\\Projetos\\Teste_neoway\\arquivo.csv": "base",
    }

    // Processar cada CSV e importá-los para o banco
    for filePath, tableName := range csvFiles {
        fmt.Printf("Importando %s para a tabela %s...\n", filePath, tableName)
        err := importCSVToPostgres(db, filePath, tableName)
        if err != nil {
            log.Fatalf("Erro ao importar %s: %v", filePath, err)
        }
        fmt.Printf("Importação de %s concluída.\n", filePath)
    }

    fmt.Println("Todas as importações foram concluídas!")
}

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Conexão com o banco de dados PostgreSQL
	connStr := "user=postgres dbname=postgres password=Mepf@2002 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificando a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Criando a tabela, se não existir
	createTableSQL := 
	CREATE base (
		id SERIAL PRIMARY KEY,
		CPF TEXT
	);
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	// Abrindo o arquivo .txt
	file, err := os.Open("seu_arquivo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Lendo o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Inserindo a linha no banco de dados
		_, err = db.Exec("INSERT INTO sua_tabela (sua_coluna) VALUES ($1)", line)
		if err != nil {
			log.Printf("Erro ao inserir: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dados armazenados com sucesso!")
}
*/

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Conexão com o banco de dados PostgreSQL
	connStr := "user=postgres dbname=postgres password=Mepf@2002 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificando a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Criando a tabela, se não existir
	createTableSQL := `
	CREATE TABLE clientes1 (
		cpf VARCHAR(100),
		private INT,
		incompleto INT,
		data_ultima_compra DATE,
		ticket_medio NUMERIC(10, 2),
		ticket_ultima_compra NUMERIC(10, 2),
		loja_frequente VARCHAR(100),
		loja_ultima_compra VARCHAR(100)
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	// Abrindo o arquivo .txt
	file, err := os.Open("C:\\Users\\franc\\Downloads\\base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Lendo o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	firstLine := true // Flag para ignorar a primeira linha (cabeçalho)
	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			firstLine = false // Ignora a primeira linha
			continue
		}

		// Dividindo a linha em colunas
		columns := strings.Fields(line)

		// Verificando se a linha tem o número correto de colunas
		if len(columns) != 8 {
			log.Printf("Linha com número de colunas inválido: %s", line)
			continue
		}

		// Convertendo os tickets para decimal
		ticketMedio, err := convertToDecimal(nullIf(columns[4]))
		if err != nil {
			log.Printf("Erro ao converter ticket_medio: %v", err)
		}

		ticketUltimaCompra, err := convertToDecimal(nullIf(columns[5]))
		if err != nil {
			log.Printf("Erro ao converter ticket_ultima_compra: %v", err)
		}

		// Inserindo os dados no banco de dados
		_, err = db.Exec(`
			INSERT INTO clientes1 (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_frequente, loja_ultima_compra) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			nullIf(columns[0]),                     // CPF
			nullIf(columns[1]),                     // PRIVATE
			nullIf(columns[2]),                     // INCOMPLETO
			nullIf(columns[3]),                     // DATA DA ÚLTIMA COMPRA
			ticketMedio,                           // TICKET MÉDIO
			ticketUltimaCompra,                    // TICKET DA ÚLTIMA COMPRA
			nullIf(columns[6]),                     // LOJA MAIS FREQUÊNTE
			nullIf(columns[7]),                     // LOJA DA ÚLTIMA COMPRA
		)
		if err != nil {
			log.Printf("Erro ao inserir: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dados armazenados com sucesso!")
}

// Função para tratar NULLs
func nullIf(s string) interface{} {
	if strings.ToUpper(s) == "NULL" || s == "" {
		return nil
	}
	return s
}

// Função para converter valores monetários de string para decimal
func convertToDecimal(value interface{}) (interface{}, error) {
	if value == nil {
		return nil, nil // Retorna nil se o valor for NULL
	}

	// Converte o valor para string
	strValue, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("valor não é uma string")
	}

	// Remove a vírgula e substitui por ponto
	strValue = strings.ReplaceAll(strValue, ",", ".")
	return strconv.ParseFloat(strValue, 64)
}
