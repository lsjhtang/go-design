package lib

import (
	"github.com/streadway/amqp"
	"log"
	"wserver/initalize"
)

const (
	QUEUE_USER         = "user"
	QUEUE_USER_PARTNER = "user_partner"
	EXCHANGE_USER      = "userExchange"
	ROUTER_KEY_PARTNER = "partner"
)

type MQ struct {
	Channel *amqp.Channel
}

func NewMQ() *MQ {
	c, err := initalize.GetMq().Channel()
	if err != nil {
		log.Panic(err)
	}
	return &MQ{Channel: c}
}

func (this *MQ) DceQueueAadBind(queues []string, key string, exchange string) error {
	for _, queue := range queues {
		q, err := this.Channel.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return err
		}

		err = this.Channel.QueueBind(q.Name, key, exchange, false, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *MQ) SendMessage(exchange string, key string, message string) error {
	err := this.Channel.Publish(exchange, key, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}

func (this *MQ) Consume(queue string, key string, callback func(<-chan amqp.Delivery)) {
	msgs, err := this.Channel.Consume(queue, key, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error:%s", err.Error())
	}

	callback(msgs)
}
