package beerstyle

import amqp "github.com/rabbitmq/amqp091-go"

type Rabbit interface {
	Publish(queue string, body []byte) error
	Consume(queue string) (<-chan amqp.Delivery, error)
}
