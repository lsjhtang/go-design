package lib

import "fmt"

func UserInit() error {

	mq := NewMQ()
	if mq == nil {
		return fmt.Errorf("MQ init error")
	}
	defer mq.Channel.Close()

	//普通队列
	//err := mq.Channel.ExchangeDeclare(EXCHANGE_USER, "direct", false, false, false, false, nil)
	//延时队列
	err := mq.Channel.ExchangeDeclare(EXCHANGE_USER_DELAY, "x-delayed-message", false, false, false, false, map[string]interface{}{"x-delayed-type": "direct"})
	if err != nil {
		return fmt.Errorf("Exchange error", err)
	}

	qs := []string{QUEUE_USER, QUEUE_USER_PARTNER}
	//err = mq.DceQueueAadBind(qs, ROUTER_KEY_PARTNER, EXCHANGE_USER)
	err = mq.DceQueueAadBind(qs, ROUTER_KEY_PARTNER, EXCHANGE_USER_DELAY)
	if err != nil {
		return fmt.Errorf("queue bind error", err)
	}
	return nil
}
