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

	// Criaar a tabela
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

	// leitura do arquivo .txt
	file, err := os.Open("C:\\Users\\franc\\Downloads\\base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	
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
