package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-cqrs-api/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

const queueName = "user_created"

func InitRabbitMQ() {
	var err error
	rabbitConn, err = amqp.Dial(config.AppConfig.RabbitURL)
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}

	rabbitChannel, err = rabbitConn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ channel error:", err)
	}

	_, err = rabbitChannel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal("Queue declaration error:", err)
	}
}

func PublishUserCreatedEvent(data map[string]string) error {
	body, _ := json.Marshal(data)

	return rabbitChannel.PublishWithContext(
		context.Background(),
		"",        // default exchange
		queueName, // routing key (queue name)
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func StartEventSubscribers() {
	msgs, err := rabbitChannel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to consume RabbitMQ messages:", err)
	}

	go func() {
		for msg := range msgs {
			var data map[string]string
			json.Unmarshal(msg.Body, &data)

			key := fmt.Sprintf("user:%s", data["id"])
			RedisClient().Set(context.Background(), key, msg.Body, 0)
			log.Println("Read model updated for user:", data["id"])
		}
	}()
}
