package configs

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDb(configs *AppConfig) (*mongo.Client, error) {

	clientOpts := options.Client().ApplyURI(configs.DOCUMENT_DB_DSN)

	if configs.ENVIRONMENT != "local" {
		caCert, err := os.ReadFile("global-bundle.pem")
		if err != nil {
			log.Fatalf("Erro ao carregar o certificado CA: %v", err)
		}

		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			RootCAs: certPool,
		}
		clientOpts.SetTLSConfig(tlsConfig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Erro ao conectar ao DocumentDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao verificar a conexão: %v", err)
	}

	log.Println("Conectado ao DocumentDB com sucesso!")
	return client, nil
}
