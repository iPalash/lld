package publisher

import (
	"errors"
	"kafka/message"
	"sync"
)

var (
	ErrorTopicAlreadyExists = errors.New("topic already exists")
	ErrorTopicDoesNotExist  = errors.New("topic does not exists")
	ErrorOffsetOverflow     = errors.New("offset too high")
)

// This should be dumb
type Publisher interface {
	CreateTopic(string) error
	Publish(string, message.Message) error
	Size(string) (int, error)
	Get(string, int) (message.Message, error)
}

type PublisherImpl struct {
	receiver sync.Map
	sizes    sync.Map
	blocked  sync.Map
	data     map[string][]message.Message
}

func New() Publisher {
	data := make(map[string][]message.Message)
	return &PublisherImpl{data: data}
}

func (p *PublisherImpl) checkTopic(topic string) (chan message.Message, bool) {
	val, ok := p.receiver.Load(topic)
	if !ok {
		return nil, ok
	} else {
		return val.(chan message.Message), ok
	}
}

func (p *PublisherImpl) CreateTopic(topic string) error {
	if _, ok := p.checkTopic(topic); ok {
		return ErrorTopicAlreadyExists
	} else {
		p.receiver.Store(topic, make(chan message.Message))
		p.sizes.Store(topic, 0)
		p.receive(topic)
		// p.data[topic] = []message.Message{}
		return nil
	}
}

func (p *PublisherImpl) Publish(topic string, m message.Message) error {
	if val, ok := p.checkTopic(topic); ok {
		val <- m
		// p.data[topic] = append(p.data[topic], m)
		return nil
	} else {
		return ErrorTopicDoesNotExist
	}
}

func (p *PublisherImpl) Size(topic string) (int, error) {
	if _, ok := p.checkTopic(topic); ok {
		sz, _ := p.sizes.Load(topic)
		return sz.(int), nil
	} else {
		return -1, ErrorTopicDoesNotExist
	}
}

func (p *PublisherImpl) Get(topic string, offset int) (message.Message, error) {
	if _, ok := p.checkTopic(topic); ok {
		for {
			sz, _ := p.sizes.Load(topic)
			if sz.(int) > offset {
				return p.data[topic][offset], nil
			} else {
				// Block this until next value hasn't come
				val, _ := p.blocked.Load(topic)
				wait := make(chan struct{})
				if val, ok := val.([](chan struct{})); ok {
					p.blocked.Store(topic, append(val, wait))
				} else {
					p.blocked.Store(topic, [](chan struct{}){wait})
				}

				<-wait // Wait on this
			}
		}
	} else {
		return message.Message{}, ErrorTopicDoesNotExist
	}
}

func (p *PublisherImpl) receive(topic string) {
	go func(t string) {
		for {
			if v, ok := p.checkTopic(t); ok {
				msg := <-v
				p.data[t] = append(p.data[t], msg)

				sz, _ := p.sizes.Load(topic)
				p.sizes.Store(topic, sz.(int)+1)
				blockers, _ := p.blocked.Load(topic)
				for blockers, ok := blockers.([]interface{}); ok; {
					for _, block := range blockers {
						(block.(chan struct{})) <- struct{}{}
					}
				}
				p.blocked.Delete(topic)
			}
		}
	}(topic)
}
