# api-golang

### Running the database:
`docker-compose up -d`

### Running with live reload:
`air`

### Running without live reload:
`go run ./cmd/main.go`


### Main dependencies:

 github.com/dgrijalva/jwt-go

 github.com/go-playground/validator/v10

 go.mongodb.org/mongo-driver/mongo

 github.com/gorilla/mux

 github.com/joho/godotenv

 github.com/stretchr/testify

 github.com/rs/cors

 github.com/stretchr/testify/mock

Air - live reload tool
go install github.com/air-verse/air@latest

K6 - load testing tool
sudo apt-get update && \\
sudo apt-get install -y gnupg2 && \\
curl -s https://dl.k6.io/key.gpg | sudo gpg --dearmor -o /usr/share/keyrings/k6-archive-keyring.gpg && \\
echo 'deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main' | sudo tee /etc/apt/sources.list.d/k6.list > /dev/null && \\
sudo apt-get update && \\
sudo apt-get install -y k6

TODO:
[ ] Review error messages
[x] English only
[x] Review whole project
[ ] Teste e2e
[x] Load test
[x] Makefile (list of all available commands)
[ ] Graceful shutdown
[ ] Use gin instead of gorilla mux
[ ] Logs, traces and metrics
[ ] Pipeline CI/CD
[ ] Email test server in docker-compose (eg. mailslurper)
[x] Vertical slicing architecture - organize folders by domain not by function