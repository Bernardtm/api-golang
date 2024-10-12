# Etapa de build
FROM golang:nanoserver-ltsc2022 AS builder

# Configura o diretório de trabalho dentro do contêiner
WORKDIR /app

RUN go install github.com/air-verse/air@latest

# Copia os arquivos de dependência do Go e faz o download dos módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código da aplicação
COPY . .

# Compila a aplicação
RUN go build -o main ./cmd/main.go

# Etapa final
FROM alpine:3.18

# Cria um usuário não-root para executar a aplicação (por segurança)
RUN adduser -D -g '' appuser

# Define o diretório de trabalho
WORKDIR /app

# Copia o binário compilado da etapa de build
COPY --from=builder /app/main .

# Ajusta as permissões
RUN chown -R appuser /app

# Muda para o usuário não-root
USER appuser

# Expõe a porta que a aplicação usará
EXPOSE 8080

# Comando para executar a aplicação
CMD ["air", "-c", ".air.toml"]