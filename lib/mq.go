package lib

import (
	"github.com/streadway/amqp"
	"log"
	"wserver/initalize"
)

const (
	QUEUE_USER          = "user"
	QUEUE_USER_PARTNER  = "user_partner"
	EXCHANGE_USER       = "userExchange"
	EXCHANGE_USER_DELAY = "userDelayExchange"
	ROUTER_KEY_PARTNER  = "partner"
)

type MQ struct {
	Channel       *amqp.Channel
	notifyConfirm chan amqp.Confirmation //确认模式
	notifyReturn  chan amqp.Return       //消息回执模式
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
	err := this.Channel.Publish(exchange, key, true, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}

//发送延时消息
func (this *MQ) SendDelayMessage(exchange string, key string, message string, ttl int) error {
	err := this.Channel.Publish(exchange, key, true, false,
		amqp.Publishing{
			Headers:     map[string]interface{}{"x-delay": ttl},
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

func (this *MQ) SetConfirm() { //消息确认模式
	err := this.Channel.Confirm(false)
	if err != nil {
		log.Print(err)
	}

	this.notifyConfirm = this.Channel.NotifyPublish(make(chan amqp.Confirmation))
}

func (this *MQ) ListenConfirm() {
	defer this.Channel.Close()
	result := <-this.notifyConfirm
	if result.Ack {
		log.Println("交换机确认成功")
	} else {
		log.Println("交换机确认失败")
	}
}

func (this *MQ) NotifyReturn() { //消息回执模式

	this.notifyReturn = this.Channel.NotifyReturn(make(chan amqp.Return))
	go this.listenReturn()
}

func (this *MQ) listenReturn() {
	result := <-this.notifyReturn
	if string(result.Body) != "" {
		log.Println("消息入队成功")
	} else {
		log.Println("消息入队失败", string(result.Body))
	}
}
