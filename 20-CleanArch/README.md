# Desafio Clean Architecture

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar o usecase de listagem das orders.

Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.


### Execução

1. Clone o repositório e acesse a subpasta do desafio:

   ```bash
   git clone https://github.com/mbombonato/pos-go-expert.git
   cd pos-go-expert/challenge-clean-architecture
   ```

2. Instale as dependências:

   ```bash
   go install github.com/ktr0731/evans@latest
   go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   go mod tidy
   ```

3. Copie as variáveis de ambiente:

   ```bash
   cp .env.sample .env
   ```

4. Execute o Mysql e RabbitMQ:

   ```bash
   docker-compose up -d
   ```

5. Execute as migrations:

   ```bash
   migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
   ```

6. Execute os servidores:

   ```bash
   go run cmd/ordersystem/wire_gen.go cmd/ordersystem/main.go
   ```

### Testando Endpoints
## 1 - API REST
Na pasta /api temos os arquivos para testes dos nossos endpoints REST na porta 8080.
- create_order.http (POST /order para criar Orders)
- list_orders.http  (GET /order para requisitar a lista de Orders)

## 2 - GraphQL
Em http://localhost:8082, podemos executar os comandos GraphQL.

Criando Orders:
```sql
mutation createOrder {
  createOrder(input: {id: "hue", Price: 10.00, Tax: 1.00}), {
    id
    Price
    Tax
    FinalPrice
  }
}
```

Listando Orders:
```sql
query listOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## 3 - gRPC
Executando o cliente gRPC:
```bash
evans -r repl -p 8081
```

Conectando-se ao package pb:
```bash
package pb
```

Conectando-se ao Service de Orders:
```bash
service OrderService
```

Criando uma Order:
```bash
call CreateOrder
```

Listando as Orders:
```bash
call ListOrders
```

## 4 - Extra Commands
Wire
```bash
wire ./cmd/ordersystem
```

GraphQL
```bash
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate
```

Protofiles
```bash
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```