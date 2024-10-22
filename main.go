package main

import (
	"bufio"        // Fornece funções para leitura de texto de entrada
	"database/sql" // Biblioteca padrão para trabalhar com bancos de dados SQL
	"fmt"          // Usado para formatação e saída
	"log"          // Para exibir mensagens de erro
	"os"           // Para interagir com o sistema operacional, como abrir arquivos
	"strconv"      // Para converter strings para números
	"strings"      // Para manipulação de strings

	_ "github.com/lib/pq" // Driver para conectar ao PostgreSQL
)

func main() {
	// Conexão com o banco de dados PostgreSQL
	connStr := "user=postgres dbname=postgres password=Mepf@2002 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//diferença entre Open e Ping
	//sql.Open: Prepara a conexão, mas não garante que está funcionando ou acessível.
	//db.Ping: Testa se a conexão está de fato funcionando e pode ser usada, disponibilidade do banco de dados.

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Cria a tabela conforme as colunas especificadas e executa o comando usando SQL no codigo
	createTableSQL := `
	CREATE TABLE clientes (
		cpf VARCHAR(100),
		private INT,
		incompleto INT,
		data_ultima_compra DATE,
		ticket_medio NUMERIC(10, 2),
		ticket_ultima_compra NUMERIC(10, 2),
		loja_frequente VARCHAR(100),
		loja_ultima_compra VARCHAR(100)
	);`

	//O símbolo _ é usado em Go para ignorar um valor de retorno de uma função que não será utilizado.
	//o caso de db.Exec, a função retorna dois valores: um resultado (sql.Result) e um erro (error). Se você não precisa do primeiro valor (o resultado da execução), você pode usar _ para ignorá-lo.
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	// leitura do arquivo .txt
	file, err := os.Open("C:\\Users\\franc\\Downloads\\base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//bufio.NewScanner: Apenas cria o scanner. Ele prepara o scanner para ler dados de uma fonte.
	//for scanner.Scan(): É o loop que percorre as linhas da entrada fornecida ao scanner e lê cada uma delas.

	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			firstLine = false
			continue
		}

		// FistLine flag para ignorar a primeira linha (cabeçalho)

		// Dividi a linha em colunas tendo o espaço como delimitador
		columns := strings.Fields(line)

		// Verifica se a linha tem o 8 colunas
		if len(columns) != 8 {
			log.Printf("Linha com número de colunas inválido: %s", line)
			continue
		}

		// Converti os tickets para decimal, para conseguir criar a tabela com esses dados.
		ticketMedio, err := convertToDecimal(nullIf(columns[4]))
		if err != nil {
			log.Printf("Erro ao converter ticket_medio: %v", err)
		}

		ticketUltimaCompra, err := convertToDecimal(nullIf(columns[5]))
		if err != nil {
			log.Printf("Erro ao converter ticket_ultima_compra: %v", err)
		}

		// Insere os dados na tabela criada de banco de dados, realizo tratamenento dos Nulls para todas as colunas
		_, err = db.Exec(`
			INSERT INTO clientes1 (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_frequente, loja_ultima_compra) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			nullIf(columns[0]), // CPF
			nullIf(columns[1]), // PRIVATE
			nullIf(columns[2]), // INCOMPLETO
			nullIf(columns[3]), // DATA DA ÚLTIMA COMPRA
			ticketMedio,        // TICKET MÉDIO
			ticketUltimaCompra, // TICKET DA ÚLTIMA COMPRA
			nullIf(columns[6]), // LOJA MAIS FREQUÊNTE
			nullIf(columns[7]), // LOJA DA ÚLTIMA COMPRA
		)
		if err != nil {
			log.Printf("Erro ao inserir: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dados armazenados!")
}

//Go não depende da ordem em que as funções são definidas no arquivo.
//Isso significa que o compilador de Go consegue localizar a definição da função, mesmo que ela esteja definida após o ponto onde foi chamada.

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
