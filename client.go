package main

import (
	"fmt"
	"log"
	"wserver/initalize"
)

func main() {
	conn := initalize.GetMq()
	defer conn.Close()

	cha, err := conn.Channel()
	if err != nil {
		log.Fatalf("error:%s", err.Error())
	}
	defer cha.Close()
	msgs, err := cha.Consume("users", "abc", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error:%s", err.Error())
	}

	for msg := range msgs {
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
