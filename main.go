/*package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abra o arquivo
	inputFile, err := os.Open("C:\\Users\\franc\\Downloads\\base_teste.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer inputFile.Close()

	// Criar o arquivo de saída (arquivo .csv)
	outputFile, err := os.Create("arquivo.csv")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo CSV:", err)
		return
	}
	defer outputFile.Close()

	// Criar um escritor CSV
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Criar um scanner para ler o arquivo .txt linha por linha
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		// Ler a linha e dividir os valores por espaços
		linha := scanner.Text()
		valores := strings.Fields(linha) // Divide por espaços em branco

		// Escrever os valores no arquivo CSV
		if err := writer.Write(valores); err != nil {
			fmt.Println("Erro ao escrever os valores:", err)
			return
		}
	}

	// Verificar se houve erros na leitura do arquivo .txt
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura do arquivo .txt:", err)
	}

	{
		// Criar um leitor CSV
		reader := csv.NewReader(outputFile)

		// Ler todo o conteúdo do arquivo CSV
		linhas, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Erro ao ler o arquivo CSV:", err)
			return
		}

		// Percorrer e imprimir cada linha do CSV
		for _, linha := range linhas {
			fmt.Println(linha)
		}

	}
}*/

package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "strings"
)

func main() {
    // Converte o arquivo .txt para .csv
    inputFile, _ := os.Open("C:\\Users\\franc\\Downloads\\base_teste.txt")
    defer inputFile.Close()

    outputFile, _ := os.Create("arquivo.csv")
    defer outputFile.Close()

    writer := csv.NewWriter(outputFile)
    
    // Lê o arquivo .txt e escreve no .csv
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        linha := scanner.Text()
        valores := strings.Fields(linha)
        writer.Write(valores)
    }
    writer.Flush()
	println("Gerado o arquivo.csv")
}

