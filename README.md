# Go Exemplo de Injeção de Dependência com Wire

Este projeto demonstra o uso de **Injeção de Dependência** em Go utilizando o framework [Google Wire](https://github.com/google/wire). O exemplo inclui a implementação de um repositório de produtos, uma camada de caso de uso (use case), e uma aplicação que interage com um banco de dados SQLite.

Para fins de demonstração, o dado não é persistido nem recuperado do SQLite. Em vez disso, a função `GetProduct` retorna um resultado estático. Isso é feito apenas para ilustrar a injeção de dependência sem a complexidade de operações reais de banco de dados.


## Estrutura do Projeto

O projeto está organizado da seguinte maneira:

- **`main.go`**: Ponto de entrada da aplicação. Configura o banco de dados SQLite e executa o caso de uso para obter informações sobre um produto.
- **`product/`**: Diretório que contém a lógica de domínio relacionada aos produtos:
  - **`product.go`**: Define a estrutura `Product`, que representa a entidade Produto.
  - **`repository.go`**: Define a interface `ProductRepositoryInterface` e sua implementação `ProductRepository`, responsável por interagir com o banco de dados.
  - **`usecase.go`**: Contém a camada de lógica de negócios `ProductUseCase`, que utiliza o repositório para recuperar informações sobre produtos.
- **`wire.go`**: Arquivo que define as dependências e configura a injeção de dependência utilizando o Google Wire.
- **`wire_gen.go`**: Arquivo gerado automaticamente pelo Wire que contém a implementação das injeções de dependência definidas em `wire.go`. Este arquivo não deve ser editado manualmente, pois é gerado automaticamente pelo comando `wire`.



## Pré-requisitos

Antes de executar o projeto, certifique-se de que seu ambiente possui as seguintes ferramentas instaladas:

- [Go](https://golang.org/dl/) (versão 1.20 ou superior)
- [SQLite3](https://www.sqlite.org/download.html)
- [Google Wire](https://github.com/google/wire)

## Instalação

Siga os passos abaixo para configurar o projeto localmente:

1. Clone o repositório:

   ```bash
   git clone https://github.com/ramonamorim/go-di.git
   cd go-di
   ````

2. Instale o wire:
    ```bash
    go install github.com/google/wire/cmd/wire@latest
    ````

3. Instale as dependências do projeto usando go mod:
    ```bash
    go mod tidy
    ```

4. Gere o código de injeção de dependências com o Wire. Navegue até o diretório que contém o arquivo `wire.go` e execute o seguinte comando:

    ```bash
    wire
    ```

## Executando o Projeto

Para executar a aplicação, siga os passos abaixo:

1. Execute a aplicação:
    ```bash
    go run main.go wire_gen.go
    ```


2. O resultado esperado será:
    ```yaml
    ID: 1, Name: Teste
    ```

## Compilação e Execução

1. Compilação:

O Go compila automaticamente o código quando você usa go run. Para uma abordagem mais tradicional, você pode compilar o código em um binário executável usando o comando go build:

```bash
go build -o app
```

2. Execução

Isso criará um arquivo executável chamado myapp (ou o nome que você escolher) no diretório atual. Você pode então executar o binário diretamente:

    ```bash
    ./app
    ```