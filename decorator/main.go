// main.go
package main

import (
	"github.com/RicardoKim/Study/DesignPattern/Decorator/decorators"
	"github.com/RicardoKim/Study/DesignPattern/Decorator/sender"
)

func main() {
	message := "Hello, Decorator Pattern!"

	sms := &sender.SimpleMessageSender{}
	encrypted := decorators.NewEncryptDecorator(sms)
	encryptedAndCompressed := decorators.NewCompressDecorator(encrypted)

	// 기본 메시지 전송
	sms.Send(message)

	// 암호화된 메시지 전송
	encrypted.Send(message)

	// 암호화되고 압축된 메시지 전송
	encryptedAndCompressed.Send(message)
}
