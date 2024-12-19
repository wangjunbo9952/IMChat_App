package kafka

import (
	"IMChat_App/pkg/protocol"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
)

var producer *kafka.Producer

func GetProducer() *kafka.Producer {
	return producer
}

func InitProducer() {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Println("Failed to create producer: %s", err)
	}
}

func SendMessage(producer *kafka.Producer, topic string, chatMessage *protocol.Message) {
	// 将结构体序列化为JSON
	messageBytes, err := proto.Marshal(chatMessage)
	if err != nil {
		fmt.Println("Failed to serialize message: %s", err)
		return
	}

	// 发送消息
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}, nil)

	if err != nil {
		fmt.Println("Failed to produce message: %s", err)
	}
}
