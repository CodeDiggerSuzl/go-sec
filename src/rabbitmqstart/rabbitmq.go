package rabbitmqstart

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// simple mode
// account:passwd@addr:port/vhost
const MqURL = "amqp://suzl:rabbit@127.0.0.1:5672/go-sec"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	MqUrl     string
}

// Create rabbit mq instance.
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitMq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: MqURL}
	var err error
	// 创建连接
	rabbitMq.conn, err = amqp.Dial(rabbitMq.MqUrl)
	rabbitMq.failOnErr(err, "error during NewSimpleRabbitMq: Error during get rabbit mq connection.")
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	rabbitMq.failOnErr(err, "Error during get channel")
	return rabbitMq
}

// Disconnect with the rabbit mq connection.
func (r *RabbitMQ) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// Error handle func.
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 1. Simple mode.
func NewSimpleRabbitMq(queueName string) *RabbitMQ {
	// Using default exchange and key def exchange.
	return NewRabbitMQ(queueName, "", "")
}

// Send msg of simple mode
func (r *RabbitMQ) SimplePublish(message string) {
	// 1. apply for queue,if no queue, will create ,if there is queue
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		log.Println("err during QueueDeclare: ", err)
	}
	// send message
	// if mandatory is true: if can't find the queue by exchange and route-key the message will return to the sender
	_ = r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}

func (r *RabbitMQ) SimpleConsume() {
	// 1. declare queue
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		log.Println("err during QueueDeclare: ", err)
	}
	// 2. consume the message
	message, err := r.channel.Consume(r.QueueName, "", true, false, false, false, nil)
	if err != nil {
		log.Println("err during message meg", err)
	}

	// 3. use channel to handle the message
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received the message: %s ", d.Body)
			// handle the message
			fmt.Println(string(d.Body))
		}
	}()
	fmt.Println("wait for message to exit")
	<-forever
}

// pub/sub mode
func NewRabbitMqPubSub(exchangeName string) *RabbitMQ {
	// create rabbit mq instance
	rabbitMQ := NewRabbitMQ("", exchangeName, "")
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.MqUrl)
	rabbitMQ.failOnErr(err, "error during create rabbit mq conn")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "err during get channel")
	return rabbitMQ
}

// publishing in publish mode
func (r *RabbitMQ) PublishPub(message string) {
	// try to create the exchange
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.failOnErr(err, "Fail to declare an exchange")
	err = r.channel.Publish(r.Exchange, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	r.failOnErr(err, "error during publish message in pub/sub mode")
}

// consume mode in pub/sub mode
func (r *RabbitMQ) ReceiveSub() {
	// try to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange, "fanout", true, false, false, false, nil)
	r.failOnErr(err, "error during rece message in pub/sub message")

	queue, err := r.channel.QueueDeclare("", false, false, true, false, nil)
	r.failOnErr(err, "error during declare a queue")

	// bind queue to exchange
	err = r.channel.QueueBind(queue.Name, "", r.Exchange, false, nil)
	// consume msg
	msg, err := r.channel.Consume(queue.Name, "", true, false, false, false, nil)

	exitChan := make(chan bool)

	go func() {
		for m := range msg {
			log.Printf("Get messge %s:", m.Body)
			fmt.Println(string(m.Body))
		}
	}()
	<-exitChan
}

// routing mode
func NewRoutingRabbitMQ(exchangeName string, routingKey string) *RabbitMQ {
	mq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	mq.conn, err = amqp.Dial(mq.MqUrl)
	mq.failOnErr(err, "error during connect to rabbit mq")

	mq.channel, err = mq.conn.Channel()
	mq.failOnErr(err, "err to open a channel")
	return mq
}

func (r *RabbitMQ) PublishingRouting(message string) {
	// type to direct
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)
	r.failOnErr(err, "err during exchange declare in publishing routing")

	// send msg routing key
	err = r.channel.Publish(r.Exchange, r.Key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

}

func (r *RabbitMQ) RecvRouting() {
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)
	r.failOnErr(err, "error during declare an exchange")
	queue, err := r.channel.QueueDeclare("", false, false, true, false, nil)

	// bind queue
	err = r.channel.QueueBind(queue.Name, r.Key, r.Exchange, false, nil)
	msg, err := r.channel.Consume(queue.Name, "", true, false, false, false, nil)

	exitChan := make(chan bool)

	go func() {
		for m := range msg {
			log.Printf("get msg in routing recv: %s", m.Body)
			fmt.Println(string(m.Body))
		}
	}()
	<-exitChan
}
