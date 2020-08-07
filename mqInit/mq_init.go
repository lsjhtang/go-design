package mqInit

import (
	"github.com/streadway/amqp"
	"log"
)

var MqConn *amqp.Connection

func init() {
	conn, err := amqp.Dial("amqp://admin:123@192.168.10.240:5672/")
	if err != nil {
		log.Fatalf("%s: %s", err, "Failed to connect to RabbitMQ")
	}
	MqConn = conn
}

func GetMq() *amqp.Connection {
	return MqConn
}
