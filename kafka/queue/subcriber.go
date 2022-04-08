package queue

import (
	"fmt"
	"math/rand"
)

type Subscriber interface {
	Start()
	Notify()
}

type SubscriberImpl struct {
	uuid   int
	offset int
	topic  Topic
	block  chan struct{}
}

func NewSub(topic Topic) Subscriber {
	block := make(chan struct{}, 1)
	sub := &SubscriberImpl{
		rand.Intn(100), 0, topic, block,
	}
	topic.AddSubscription(sub)
	return sub
}

func (s *SubscriberImpl) Start() {
	go func() {
		for {

			for s.topic.Size() > s.offset {
				s.process(s.topic.Get(s.offset))
				s.offset += 1
			}
			<-s.block
			// time.Sleep(time.Minute)
			// else wait

		}
	}()
}

func (s *SubscriberImpl) process(msg Message) {
	println(fmt.Sprintf("%d: processed num:%d => %v", s.uuid, s.offset, msg))
}

func (s *SubscriberImpl) Notify() {
	s.block <- struct{}{}
}
