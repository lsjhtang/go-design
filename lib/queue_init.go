package lib

import "fmt"

func UserInit() error {

	mq := NewMQ()
	if mq == nil {
		return fmt.Errorf("MQ init error")
	}
	defer mq.Channel.Close()

	err := mq.Channel.ExchangeDeclare("userExchange", "direct", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Exchange error", err)
	}

	qs := []string{QUEUE_USER, QUEUE_USER_PARTNER}
	err = mq.DceQueueAanBind(qs, ROUTER_KEY_PARTNER, EXCHANGE_USER)
	if err != nil {
		return fmt.Errorf("queue bind error", err)
	}
	return nil
}
