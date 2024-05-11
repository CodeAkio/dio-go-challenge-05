# Calculadora Básica em Go

Este projeto consiste em uma simples calculadora escrita em Go que suporta operações básicas como soma, subtração, multiplicação e divisão. O projeto também inclui testes unitários para validar as funcionalidades da calculadora.

## Descrição

A calculadora oferece as seguintes funções:
- Adição (`add`)
- Subtração (`subtract`)
- Multiplicação (`multiply`)
- Divisão (`divide`)

Cada função é projetada para realizar sua operação correspondente e retornar o resultado. A função de divisão também trata o caso de divisão por zero, retornando um erro.

## Como Executar

Para executar o programa principal que demonstra as operações da calculadora, siga estas instruções:

1. Clone este repositório ou baixe os arquivos para o seu computador.
2. Abra um terminal.
3. Navegue até o diretório onde os arquivos estão salvos.
4. Execute o comando `go run main.go`.

## Como Executar os Testes

Para executar os testes unitários associados às funções da calculadora, siga estas etapas:

1. Certifique-se de estar no diretório do projeto onde os arquivos `main.go` e `main_test.go` estão localizados.
2. No terminal, execute o comando `go test`.

Os testes verificarão se as operações matemáticas estão sendo realizadas corretamente e se a divisão por zero está sendo tratada adequadamente.
