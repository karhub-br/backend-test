package rabbit

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(host, user, pass, port string) (*amqp.Connection, *amqp.Channel) {
	rabbitInfo := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		user, pass, host, port)

	conn, err := amqp.Dial(rabbitInfo)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Erro ao abrir canal:", err)
	}
	return conn, ch
}
