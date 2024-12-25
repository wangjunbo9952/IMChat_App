package rabbitmq

import "IMChat_App/internal/mq/rabbit"

// RabbitMQ的SendMsg API
func SendMsgToMq(msg []byte) error {
	err := rabbit.SendMsgToMq(msg)
	if err != nil {
		return err
	}
	return nil
}
