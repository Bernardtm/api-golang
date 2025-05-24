package queues

import (
	"testing"

	"bernardtm/backend/pkg/redis_client"
)

type TestMessage struct {
	Task string `json:"task"`
}

func TestEnqueue(t *testing.T) {
	// Create a new Redis client
	redisClient, err := redis_client.ConnectRedis("localhost:6379")
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Create a new RedisQueueProvider
	queueProvider := NewRedisQueueProvider(redisClient)

	// Enqueue a message
	queueProvider.Enqueue(TestMessage{Task: "test-task"}, "task_stream")
}
