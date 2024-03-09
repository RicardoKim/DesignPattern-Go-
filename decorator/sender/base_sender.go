package sender

import (
	"fmt"
)

type SimpleMessageSender struct{}

func (sms *SimpleMessageSender) Send(message string) error {
	fmt.Println("Sending message:", message)
	return nil
}
