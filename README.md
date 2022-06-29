##  Desafio Q2Bank:
- Temos 2 tipos de usuários, os comuns e lojistas, ambos têm carteira com dinheiro e realizam transferências entre eles. Vamos nos atentar somente ao fluxo de transferência entre dois usuários.

## Requisitos:

Essa aplicação requer a versão 1.18.2+ do [Golang]((https://go.dev/doc/install) instalado e também do [Docker Compose](https://docs.docker.com/compose/install/).
Instale as dependências e para rodar a aplicação use o passo-a-passo abaixo:

### Passo 1:
Suba os containers rodando o comando abaixo na pasta raiz do projeto:
```sh
docker-compose up
```

### Passo 2:
Após a inicialização do docker-compose, ainda na raiz e em um novo terminal, rode o comando:
```sh
go run router.go
```

- *Após efetuar esses passos, sua aplicação deverá estar no ar!*

# Métodos
Requisições para a API devem seguir os padrões:
| Método | Descrição |
|---|---|
| `POST` | Insere uma nova transação no Banco de Dados|

## Notas

Deixei alguns Usuários já criados, é possivel visualizá-los no Banco de Dados Postgres para efeito de testes, e o schema na pasta cmd/schema.sql.

Endpoint para testar as transações - http://localhost:8080/transaction
Payload: 
```{
    "value": ,
    "payee": ,
    "payer": 
}```

Utilizei uma arquitetura de monolito pois é algo que já estava mais habituado em codificar, tive alguns problemas durante o desenvolvimento por estar utilizando ambiente Windows e com me demandou mais tempo para a entrega do desafio. Esse desafio da Q2Bank de fato foi um desafio, por nunca ter usado o Postgres e também em um ambiente docker no Windows.
