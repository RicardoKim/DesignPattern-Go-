// decorators/encrypt_decorator.go
package decorators

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/Decorator/sender"
)

type EncryptDecorator struct {
	*MessageSenderDecorator
}

func NewEncryptDecorator(ms sender.MessageSender) *EncryptDecorator {
	return &EncryptDecorator{NewMessageSenderDecorator(ms)}
}

func (ed *EncryptDecorator) Send(message string) error {
	encryptedMessage := "encrypted(" + message + ")"
	fmt.Println("Encrypting message...")
	return ed.Wrapped.Send(encryptedMessage)
}
