package sender

type MessageSender interface {
	Send(message string) error
}
