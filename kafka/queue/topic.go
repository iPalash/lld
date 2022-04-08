package queue

import (
	"sync"
)

type Topic interface {
	Add(Message)
	AddSubscription(Subscriber)
	Size() int
	Get(int) Message
}

type TopicImpl struct {
	name string
	lock *sync.RWMutex
	Msgs []Message
	subs []Subscriber
}

func NewTopic(s string) Topic {
	return &TopicImpl{s, &sync.RWMutex{}, []Message{}, []Subscriber{}}
}

func (t *TopicImpl) Add(msg Message) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.Msgs = append(t.Msgs, msg)
	go func() {
		for _, sub := range t.subs {
			sub.Notify()
		}
	}()
}

func (t *TopicImpl) AddSubscription(s Subscriber) {
	t.subs = append(t.subs, s)
}

func (t *TopicImpl) Size() int {
	return len(t.Msgs)
}

func (t *TopicImpl) Get(offset int) Message {
	return t.Msgs[offset]
}
