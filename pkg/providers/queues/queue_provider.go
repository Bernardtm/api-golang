package queues

// QueueProvider defines the interface for a queue provider
type QueueProvider interface {
	Enqueue(message interface{}, queueName string) error
}
