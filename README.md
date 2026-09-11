# Arquitetura-hexagonal

Projeto de gestão de produtos desenvolvido durante estudo e aplicação dos fundamentos de Arquitetura Hexagonal.

Projeto conteinerizado com Dockerfile e docker compose, utilizando linguagem `go` e banco de dados `sqlite3`.

Para executar o projeto e realizar os testes, siga os comando abaixo.


## Execução do projeto
Para executar o projeto, execute no terminal:

`docker compose up -d`

## Execução de testes unitários

A execução dos testes unitários não depende do container. Caso esteja utilizando uma IDE com suporte a Go, é possível executar os testes diretamente pelo botão Play da IDE.

Também é possível executar os testes pelo terminal.

### Execução via terminal

- Para executar um único teste:

    `go test -run <nome-do-teste>`

    É necessário estar no pacote onde o teste está localizado. Caso não esteja, informe o caminho do pacote. Por exemplo:
    
    `go test ./application -run <nome-do-teste>`

- Para executar todos os testes de um pacote:

    `go test ./application`

- Para executar todos os testes do projeto:

    `go test ./...`

### Execução via container
Também é possível executar os testes dentro do container da aplicação.

Nesse caso, o projeto deve estar em execução:

`docker compose up -d`

Os mesmos comandos utilizados no terminal podem ser executados dentro do container. Por exemplo:

`docker exec -it appproduct go test ./application -run <nome-do-teste>`

## Testando adaptadores
Os testes dos adaptadores dependem do ambiente do container. Portanto, antes de executá-los, inicie o projeto:

`docker compose up -d`

A aplicação possui três adaptadores:
- DB
- CLI
- Webserver

O arquivo `main.go` é utilizado para controlar quais adaptadores serão executados. Para testar cada adaptador, é necessário comentar/descomentar o código correspondente no arquivo `main.go`.

    IMPORTANTE: os adaptadores CLI e Webserver utilizam Cobra. Por isso, ambos utilizam o mesmo trecho de código no arquivo `main.go`.

Todos os adaptadores devem persistir os produtos no SQLite3. Os dados podem ser consultados executando os seguintes comandos:

- Acesse o container:

    `docker exec -it appproduct bash`
- Acesse o banco SQLite:

    `sqlite3 db.sqlite`
- Execute uma consulta:

    `select * from products;`

### Testando DB adapter
- Comentar/descomentar as funções `main` no arquivo `main.go` de acordo com o adaptador que deseja executar
- Acessar o container de forma iterativa:

    `docker exec -it appproduct bash`

- Executar o arquivo `main.go`
    
    `go run main.go`

Após a execução, o adaptador deve realizar a operação definida no método main. O resultado pode ser consultado no banco de dados SQLite.

### Testando CLI adapter
Para executar o adaptador cCLI, foi criado o comando `cli` que recebe parâmetros para realizar as operações. 

Os parâmetros podem ser consultados utilizando o comando `--help`.

`go run main.go cli --help`

Passo a passo para testar a cli:
- Comentar/descomentar as funções `main` no arquivo `main.go` de acordo com o adaptador que deseja executar
- Acessar o container:

    `docker exec -it appproduct bash`

- Executar o comando desejado na cli:

    `go run main.go cli -a=create -n="Product CLI" -p=29.99`

Após a execução, o adaptador deve realizar a operação solicitada. O resultado será apresentado no shell e também poderá ser consultado no banco de dados.

### Testando WEBSERVER adapter
Para executar o Webserver, foi criado o comando `http`.

As rotas disponíveis podem ser consultadas no arquivo: 

`/adapters/web/handler/product.go`

Passo a passo para testar:
- Comentar/descomentar as funções `main` no arquivo `main.go` de acordo com o adaptador que deseja executar
- Acessar o container:

    `docker exec -it appproduct bash`

- Executar o comando para iniciar o webserver:
    `go run main.go http`

Após executar o shell deve informar: 

`Webserver has been started` 

O Webserver estará disponível na porta `9000`.

As requisições podem ser realizadas via curl ou uma ferramenta de testes de API, como Postman.

Exemplo utilizando `curl`: 

`curl http://localhost:9000/product/`
