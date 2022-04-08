package queue

type Message struct {
	Text string
}

func NewMessage(s string) Message {
	return Message{s}
}
