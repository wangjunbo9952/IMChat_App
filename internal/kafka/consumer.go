package kafka

import (
	"IMChat_App/internal/service"
	"IMChat_App/internal/websocket"
	"IMChat_App/pkg/protocol"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
)

var consumer *kafka.Consumer

func GetConsumer() *kafka.Consumer {
	return consumer
}

func InitConsumer(groupID string) {
	var err error
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println("Failed to create consumer: %s", err)
	}
	return
}

func ConsumeMessages(consumer *kafka.Consumer, topic string) {
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		rawMsg, err := consumer.ReadMessage(-1)
		websocket.MyServer.Broadcast <- rawMsg.Value
		if err == nil {
			msg := &protocol.Message{}
			proto.Unmarshal(rawMsg.Value, msg)
			service.SaveMessage(msg)
			service.SaveMessageToRds(rawMsg.Value, msg.From, msg.To)
		} else {
			fmt.Println("Error while consuming message: %s", err)
		}
	}
}
