package rabbitmqstart

import (
	"fmt"
	"testing"
)

func TestNewRabbitMQ(t *testing.T) {
	smpMQ := NewSimpleRabbitMq("go-sec-start")
	smpMQ.SimplePublish("Hello go sec")
	fmt.Println("send ok")
}
