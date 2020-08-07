package main

import (
	"fmt"
	"log"
	"wserver/mqInit"
)

func main() {
	conn := mqInit.GetMq()
	defer conn.Close()

	cha, err := conn.Channel()
	if err != nil {
		log.Fatalf("error:%s", err)
	}
	defer cha.Close()
	msgs, err := cha.Consume("users", "abc", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error:%s", err)
	}

	for msg := range msgs {
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
