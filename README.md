# Orders System

Aplicação desenvolvida com **Clean Architecture** em Go, expondo três interfaces de comunicação: REST, gRPC e GraphQL.

## 🏗️ Arquitetura

O projeto segue os princípios da **Clean Architecture**, organizando o código em camadas:

- **Entity**: Regras de negócio da aplicação
- **Use Case**: Casos de uso da aplicação
- **Interface**: Adaptadores de interface (controllers, gateways)
- **Infrastructure**: Frameworks e drivers (banco de dados, mensageria)

## 🚀 Tecnologias

- **Go** — Linguagem principal
- **MySQL** — Banco de dados relacional
- **RabbitMQ** — Message broker
- **REST** — API HTTP
- **gRPC** — Remote Procedure Call
- **GraphQL** — Query language para API

## 📡 Portas

| Serviço  | Porta | Protocolo |
|----------|-------|-----------|
| REST     | 8000  | HTTP      |
| gRPC     | 50051 | gRPC      |
| GraphQL  | 8080  | HTTP      |
| MySQL    | 3306  | TCP       |
| RabbitMQ | 5672  | AMQP      |
| RabbitMQ Management | 15672 | HTTP |

## 🐳 Como executar

Basta rodar o comando abaixo, todo o ambiente será provisionado automaticamente:

```bash
docker compose up --build
```

> O Docker Compose irá:
> 1. Subir o MySQL e aguardar estar saudável
> 2. Subir o RabbitMQ
> 3. Executar as migrations automaticamente
> 4. Iniciar a aplicação

## 🔌 Exemplos de uso

### REST
```bash
curl -X POST http://localhost:8000/order
curl -X GET http://localhost:8000/listOrders
```

### GraphQL
Acesse o playground em:
```
http://localhost:8080
```

### gRPC
Para testar use o Postman:
CLique em:
```
 import a proto file
```

![gRPC no Postman](http://rafael-developer.com/wp-content/uploads/2026/06/Captura-de-Tela-2026-06-05-as-16.27.46.png)

Insira seu order.proto em:
```
 internal/infra/grpc/protofiles
```
E só colocar os playload e inserir ou listar.
Payload:
```
{
"id":"rzq",
"price":12,
"tax":13
}
```
