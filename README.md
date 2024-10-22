<h1>Neoway- Teste de Aprendizagem</h1>

<p>Este projeto realiza o armazenamento de dados em txt para uma tabela de banco de dados.<p>

## Sumário
1. [Visão Geral](#visão-geral)
2. [Pré-requisitos](#pré-requisitos)
3. [Como Utilizar](#como-utilizar)
4. [Banco de dados](#banco-de-dados)
5. [Pontos de Melhoria](#pontos-de-melhoria)


## Visão Geral

O programa conecta ao banco de dados PostgreSQL, cria uma tabela utilizando os dados de um arquivo txt, realiza o tratamento necessário para o arquivo que consiga suprir a demanda e prosseguir com sua criação no banco de dados..

## Pré-requisitos
Antes de executar o projeto, certifique-se de que você tenha os seguintes pré-requisitos instalados em seu sistema:

**Go:** Certifique-se de ter o Go instalado. Você pode baixar a versão mais recente em golang.org.<br>
**PostgreSQL:** O banco de dados PostgreSQL deve estar instalado e em execução. Você pode baixá-lo em postgresql.org.<br>
**Biblioteca:**  Necessário instalação da biblioteca para conectar no PostgreSQL:
```
go get github.com/lib/pq 
```
## Como utilizar

1. Após os pré-requisitos serem realizados, salve o arquivo em seu computador.
2. Para conseguir preencher as informações conforme o registro criado para o seu banco de dados no PostgreeSQL
```
"user= seu-usuário dbname= nome-do-seu-db password= senha sslmode= modo-de-conexão"
```
3. Após isso execute o programa no console
```+++++++++++++++++++++++++++++++++++++++++
go run ola.go
```
## Banco de dados

1. Após o progrma criar e enviar os dados necessários para o PostgreSQL, realizei a criação de duas funções para realizar as validações dos campos de CPFs e CNPJs. Segue os arquivos com a criação das funções 'Função_validação_cnpj.sql' e 'Função_validação_cpf.sql'

2. Com a função criada, criei as três colunas para informar se os campos que informam CPFs e o CNPJs são validos e criei o link entre a função e as colunas de validação. Segue o arquivo com o update 'Update_validação.sql'

3. No ultimo tratamento com os dados criei um update para realizar higienização dos dados no SQL, onde retira espaçamento desncessários, acentos.
Segue o arquivo com o tratamento: 'Higienização_ dados.sql'

## Pontos de Melhoria

Alguns pontos que eu vejo de melhoria no projeto criado.
1. **Não Adaptablidade** O programa não se adapta a outros tipos de arquivo, com outras informações. Nele reecrio as colunas para conseguir condicionar as informações de acordo, caso eu tenha o mesmo tipo de arquivo com 200 colunas terei um grande retrabalho nesse ponto
2. **Senha de banco de dados** Devido o compartilhamento do programa, visando segurança, o ideal seria termos uma forma de ocultar a senha ou não permitir que ela seja visivel.



    