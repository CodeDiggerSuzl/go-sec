package main

import (
	"fmt"
	"go-sec/src/rabbitmqstart"
)

func main() {
	smpMQ := rabbitmqstart.NewSimpleRabbitMq("go-sec-start")
	smpMQ.SimplePublish("Hello go sec")
	fmt.Println("Send message from simple mode provider")
}
