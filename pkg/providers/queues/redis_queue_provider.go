package queues

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisQueueProvider struct {
	Client *redis.Client
}

// Initialize a new RedisQueueProvider
func NewRedisQueueProvider(redisClient *redis.Client) *RedisQueueProvider {
	return &RedisQueueProvider{
		Client: redisClient,
	}
}

// Enqueue adds an message to the Redis queue
func (r *RedisQueueProvider) Enqueue(message interface{}, queueName string) error {
	// Serialize the message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return err
	}

	// Convert the serialized message into a map for XAdd
	messageMap := map[string]interface{}{"message": messageBytes}

	// Add the message to the queue in Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use XAdd to add a new entry to the stream
	_, err = r.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: queueName,
		Values: messageMap,
	}).Result()
	if err != nil {
		log.Printf("Error adding message to Redis stream: %v", err)
		return err
	}

	log.Printf("Message added to queue successfully")
	return nil
}
