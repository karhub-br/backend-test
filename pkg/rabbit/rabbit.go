package rabbit

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbit struct {
	channel *amqp.Channel
}

func (r *rabbit) Publish(queue string, body []byte) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return r.channel.PublishWithContext(context.Background(), "", queue, false, false,
		amqp.Publishing{ContentType: "text/plain", Body: body})
}

func (r *rabbit) Consume(queue string) (<-chan amqp.Delivery, error) {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return r.channel.Consume(queue, "", true, false, false, false, nil)
}

func NewRabbit(channel *amqp.Channel) *rabbit {
	return &rabbit{channel: channel}
}

