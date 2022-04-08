package subscriber

import (
	"fmt"
	"kafka/message"
	"kafka/publisher"
	"math/rand"
)

type Subscriber interface {
	Start()
}

type SubscriberImpl struct {
	uuid   int
	pub    publisher.Publisher
	offset int
	topic  string
}

func New(pub publisher.Publisher, topic string) Subscriber {
	return &SubscriberImpl{
		rand.Intn(100), pub, 0, topic,
	}
}

func (s *SubscriberImpl) Start() {
	go func() {
		for {
			msg, _ := s.pub.Get(s.topic, s.offset)
			s.execute(msg)
			s.offset += 1
		}
	}()
}

func (s *SubscriberImpl) execute(m message.Message) {
	fmt.Println(s.uuid, ": printing : ", m.Text)
}
