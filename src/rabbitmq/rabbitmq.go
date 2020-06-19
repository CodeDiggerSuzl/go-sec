package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

// simple mode
// 账号:密码@地址:port/vhost
const MQ_URL = "amqp://guest:guest@127.0.0.1:15672/test"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 链接信息
	MqUrl string
}

// 创建 rabbit mq 实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitMq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: MQ_URL}
	var err error
	// 创建连接
	rabbitMq.conn, err = amqp.Dial(rabbitMq.MqUrl)
	rabbitMq.failOnErr(err, "error during NewSimpleRabbitMq:创建连接错误")
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	rabbitMq.failOnErr(err, "获取 channel 失败")
	return rabbitMq
}

// 断开 channel 和链接
func (r *RabbitMQ) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		// panic(fmt.Sprintf("%s:%s", message, err))
		log.Fatalf("%s:%s", message, err)
	}
}

// 实现 simple 模式
func NewSimpleRabbitMq(queueName string) *RabbitMQ {
	// 使用默认的 Exchange 和 key
	return NewRabbitMQ(queueName, "", "")
}
