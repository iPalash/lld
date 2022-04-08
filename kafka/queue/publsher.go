package queue

import (
	"errors"
	"sync"
)

var (
	ErrorTopicAlreadyExists = errors.New("topic already exists")
	ErrorTopicDoesNotExist  = errors.New("topic does not exists")
	ErrorOffsetOverflow     = errors.New("offset too high")
)

// This should be dumb
type Publisher interface {
	CreateTopic(string) (Topic, error)
	Publish(string, Message) error
}

type PublisherImpl struct {
	lock *sync.RWMutex
	data map[string]Topic
}

func New() Publisher {
	data := make(map[string]Topic)
	return &PublisherImpl{&sync.RWMutex{}, data}
}

func (p *PublisherImpl) checkTopic(topic string) (Topic, bool) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if _, ok := p.data[topic]; ok {
		return p.data[topic], ok
	} else {
		return nil, ok
	}
}

func (p *PublisherImpl) CreateTopic(t string) (Topic, error) {
	if v, ok := p.checkTopic(t); ok {
		return v, ErrorTopicAlreadyExists
	} else {
		p.lock.Lock()
		defer p.lock.Unlock()
		newTopic := NewTopic(t)
		p.data[t] = newTopic
		return newTopic, nil
	}
}

func (p *PublisherImpl) Publish(topic string, m Message) error {
	if val, ok := p.checkTopic(topic); ok {
		val.Add(m)
		// p.data[topic] = append(p.data[topic], m)
		return nil
	} else {
		return ErrorTopicDoesNotExist
	}
}
