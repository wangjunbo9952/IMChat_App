package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
)

// 发送消息到RabbitMQ
func SendMsgToMq(msg []byte) error {
	var err error

	conn := pool.GetConnection()
	defer pool.ReturnConnection(conn)

	ch := pool.GetChannel(conn)
	defer pool.ReturnChannel(ch)

	err = ch.Publish(
		"",        // 默认交换机
		"default", // 队列名称
		false,     // 是否强制
		false,     // 是否立即
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})
	return err
}

// 消费消息并处理
func ConsumeMessages(handler func([]byte)) {
	conn := pool.GetConnection()
	defer pool.ReturnConnection(conn)

	ch := pool.GetChannel(conn)
	defer pool.ReturnChannel(ch)

	_, err := ch.QueueDeclare(
		"default", // 队列名称
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)

	if err != nil {
		fmt.Printf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		"default", // 队列名称
		"",        // 消费者名称
		true,      // 自动确认
		false,     // 独占
		false,     // 不等待
		false,     // 不阻塞
		nil,       // 额外属性
	)
	if err != nil {
		fmt.Printf("Failed to register a consumer: %s", err)
	}

	for msg := range msgs {
		handler(msg.Body) // 调用处理函数
	}

}
