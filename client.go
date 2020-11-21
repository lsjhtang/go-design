package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
	"wserver/goft"
	"wserver/lib"
)

func SendMail(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		go send(msg)
	}

}

func send(msg amqp.Delivery) {
	time.Sleep(time.Second * 1)
	fmt.Println(msg.DeliveryTag, string(msg.Body))
	//msg.Reject(false) //丢弃消息
	msg.Ack(false)
}

func main() {
	mq := lib.NewMQ()
	mq.Consume(lib.QUEUE_USER, "c1", SendMail)

	err := mq.Channel.Qos(20, 0, false) //限流
	if err != nil {
		goft.Error(err)
	}
}
