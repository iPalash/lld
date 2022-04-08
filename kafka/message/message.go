package message

type Message struct {
	Text string
}

func New(s string) Message {
	return Message{s}
}
