package redis_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// ConnectRedis initializes and returns a new Redis client.
func ConnectRedis(redisAddress string) (*redis.Client, error) {
	// Initialize the Redis client with configuration
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	// Use context to manage connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Ping the Redis server to ensure the connection is alive
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("could not connect to Redis: %w", err)
	}

	log.Println("Connected to Redis successfully")
	return client, nil
}
