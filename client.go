package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
	"wserver/lib"
)

func SendMail(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		fmt.Println(msg.DeliveryTag, string(msg.Body))
		time.Sleep(time.Second * 1)
		msg.Ack(false)
	}

}

func main() {
	mq := lib.NewMQ()
	mq.Consume(lib.QUEUE_USER, "c1", SendMail)
}
