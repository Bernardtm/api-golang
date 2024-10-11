api rest usando jwt

endpoints:
- cadastro
- login
- recuperação de senha
- listagem de usuarios

// melhores praticas, validacoes e tratamento adequado de erros

// Certifique-se de que o código esteja bem documentado e inclua instruções claras para a
// configuração e execução do projeto, incluindo dependências e rotinas de inicialização.

TODO: pagina get usuarios, testar chamada de frontend, verificando CORS

docker-compose up -d

go test ./... -coverprofile cover.out 
go tool cover -func cover.out 

ao criar usuario, verificar se ja existe email
usar uuid




// Principais dependencias:
// github.com/dgrijalva/jwt-go
// github.com/go-playground/validator/v10
// go.mongodb.org/mongo-driver/mongo
// github.com/gorilla/mux
// github.com/joho/godotenv
// github.com/stretchr/testify

go install
// github.com/air-verse/air@latest - live reload
air init


