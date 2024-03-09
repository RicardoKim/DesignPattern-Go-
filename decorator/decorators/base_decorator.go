package decorators

import "github.com/RicardoKim/Study/DesignPattern/Decorator/sender"

// MessageSenderDecorator 구조체
type MessageSenderDecorator struct {
	Wrapped sender.MessageSender
}

func NewMessageSenderDecorator(ms sender.MessageSender) *MessageSenderDecorator {
	/*
	1. Component에 해당하는 MessageSender 인터페이스 타입을 파라미터로 받는 것을 알 수 있다.
	2. MessageSenderDecorator의 새 인스턴스를 생성한다. 이 구조체는 Wrapped필드를 가지고 있으며 이 필드에 메세지 전송기능을 실제로 구현하는 인스턴스를 저장.
	*/
	return &MessageSenderDecorator{Wrapped: ms}
}

func (msd *MessageSenderDecorator) Send(message string) error {
	return msd.Wrapped.Send(message)
}
