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

## Documentação do Swagger

A documentação Swagger ainda não está disponível

### Exemplos de Uso

1. **`POST http://[HOST]/products`**
   - **Descrição**: Para registrar um novo produto, faça uma solicitação POST com os detalhes do produto no corpo da solicitação.

2. **`GET http://[HOST]/products/{productID}`**
   - **Descrição**: Faça uma solicitação GET para recuperar informações sobre um produto específico.

3. **`GET http://[HOST]/products`**
   - **Descrição**: Faça uma solicitação GET para listar todos os produtos.

4. **`PUT http://[HOST]/products/{productID}`**
   - **Descrição**: Faça uma solicitação PUT com os detalhes do produto no corpo da solicitação para atualizar as informações de um produto.

5. **`DELETE http://[HOST]/products/{productID}`**
   - **Descrição**: Faça uma solicitação DELETE para excluir um produto.

6. **`GET http://[HOST]/health`**
   - **Descrição**: Faça uma solicitação GET para verificar a saúde do sistema.
