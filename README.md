# TheX - Products API

Este é um projeto Go para gerenciar produtos dentro do ecossistema Market-TheX. Ele serve como o componente principal para o registro e manutenção de produtos. Abaixo estão os principais detalhes deste projeto:

- **Gerenciamento de Projeto:** Go Modules
- **Linguagem:** Go
- **Versão do Go:** 1.17

## Visão Geral

O projeto TheX - Products API faz parte do ecossistema Market-TheX e é responsável pelo gerenciamento de produtos. Sua principal finalidade é fornecer uma API robusta para o registro e administração de informações de produtos na plataforma.

## Dependências

Este projeto utiliza várias dependências e bibliotecas para garantir seu funcionamento tranquilo:

- **GORM:** Uma biblioteca ORM (Object-Relational Mapping) fantástica para Golang, que oferece um conjunto robusto de recursos para trabalhar com bancos de dados. [Mais informações](https://gorm.io/)

- **SQLite3:** Um poderoso mecanismo de banco de dados embutido escrito em Go. Esse driver permite interagir perfeitamente com bancos de dados SQLite em projetos Go. [Detalhes](https://pkg.go.dev/github.com/mattn/go-sqlite3)

## Como Começar

1. Clone este repositório em seu ambiente de desenvolvimento local.

2. Abra o projeto em seu ambiente de desenvolvimento Go preferido (por exemplo, Visual Studio Code, GoLand).

3. Certifique-se de ter um banco de dados SQLite configurado e ajuste as configurações de conexão em `config.go` para corresponder à configuração do seu banco de dados.

4. Execute o projeto e você pode acessar a API de Produtos através da URL: `http://localhost:8081/api/products`.

## Live Reload

Para facilitar o desenvolvimento, recomendamos a utilização do Live Reload. Para configurar o Live Reload, você pode usar a ferramenta `air`. Instale-a com o seguinte comando:

```bash
go install github.com/cosmtrek/air@latest
```


Depois de instalar o air, você pode iniciar o Live Reload com o seguinte comando:

```bash
air
```

Isso permitirá que as alterações no código sejam automaticamente refletidas no seu servidor em execução.

## Exemplos de Uso

Aqui estão alguns exemplos de uso para a API de Produtos:

1. **Registro de Produto** `http://[HOST]/products`
    - Para registrar um novo produto, faça uma solicitação POST com os detalhes do produto no corpo da solicitação.

2. **Recuperação de Informações de Produto** `http://[HOST]/products/{productID}`
    - faça uma solicitação GET para recuperar informações sobre um produto específico.

3. **Listagem de Produtos** `http://[HOST]/products`
    - faça uma solicitação GET para para listar todos os produtos.

4. **Atualização de Produto** `http://[HOST]/products/{productID}`
    - Faça uma solicitação PUT com os detalhes do produto no corpo da solicitação para atualizar as informações de um produto.

5. **Exclusão de Produto** `http://[HOST]/products/{productID}`
    - Faça uma solicitação DELETE para excluir um produto.

6. **Verificação de Saúde do Sistema** `http://[HOST]/health`
    - Faça uma solicitação GET para verificar a saúde do sistema.

Certifique-se de usar as rotas apropriadas para realizar as operações desejadas na API de Produtos.

## Executando o Projeto no Docker

### Pré-requisitos

Certifique-se de que você tenha o Docker instalado em seu sistema. Caso contrário, você pode baixá-lo e instalá-lo a partir do [site oficial do Docker](https://www.docker.com/get-started).

### Passo 1: Build da Imagem Docker

No diretório raiz do projeto, execute o seguinte comando para criar uma imagem Docker do projeto. Certifique-se de que esteja no mesmo diretório onde se encontra o arquivo `Dockerfile`.

```shell
docker build . -t api-thex-products --no-cache
```

### Passo 2: Executando o Contêiner Docker

Agora que a imagem Docker foi criada, você pode executar o projeto em um contêiner Docker com o seguinte comando:

```shell
docker run --name api-thex-products \
--network local \
-p 3002:3002 \
-e API_ADDRESS="0.0.0.0:3002" \
-d api-thex-products
```

### Passo 3: Acessando a Aplicação

Após a execução bem-sucedida do contêiner, você pode acessar a aplicação em http://localhost:3002. Certifique-se de que a porta 3002 esteja disponível em seu sistema.

### Encerrando o Contêiner

Para parar e remover o contêiner em execução, você pode usar o seguinte comando:

```shell
docker stop api-thex-products
docker rm api-thex-products
```

Isso encerrará o contêiner e o removerá do sistema.

```shell
docker tag api-thex-products miguelsmuller/api-thex-products:1.0.0
```



