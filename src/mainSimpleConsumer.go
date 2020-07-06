package main

import (
	"go-sec/src/rabbitmqstart"
)

func main() {
	smpMQ := rabbitmqstart.NewSimpleRabbitMq("go-sec-start")
	smpMQ.SimpleConsume()
}
