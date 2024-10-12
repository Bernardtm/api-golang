# api-golang

### Rodando o banco:
`docker-compose up -d`

### Rodando a aplicação com live reload:
`air`

### Rodando a aplicação sem live reload:
`go run ./cmd/main.go`


### Principais dependências:

 github.com/dgrijalva/jwt-go
 github.com/go-playground/validator/v10
 go.mongodb.org/mongo-driver/mongo
 github.com/gorilla/mux
 github.com/joho/godotenv
 github.com/stretchr/testify
 github.com/rs/cors
 github.com/stretchr/testify/mock

Air - utilizado para live reload
go install github.com/air-verse/air@latest

TODO:
[ ] Rever mensagens de retorno
[ ] Padronizar tudo em portugues
[ ] Revisar todo o projeto
[ ] Criar testes e2e

O que seria bom que esse projeto tivesse para estar em produção?
Monitoramento de logs, traces e métricas
Configuração de pipeline de deploy
Servidor de teste de email no docker-compose (ex: mailslurper)
Teste de carga
Teste e2e
