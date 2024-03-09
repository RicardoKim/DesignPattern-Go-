// decorators/compress_decorator.go
package decorators

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/Decorator/sender"
)

type CompressDecorator struct {
	*MessageSenderDecorator
}

func NewCompressDecorator(ms sender.MessageSender) *CompressDecorator {
	return &CompressDecorator{NewMessageSenderDecorator(ms)}
}

func (cd *CompressDecorator) Send(message string) error {
	compressedMessage := "compressed(" + message + ")"
	fmt.Println("Compressing message...")
	return cd.Wrapped.Send(compressedMessage)
}
