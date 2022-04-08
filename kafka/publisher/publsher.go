package publisher

import (
	"errors"
	"kafka/message"
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
	data map[string][]message.Message
}

func New() Publisher {
	data := make(map[string][]message.Message)
	return &PublisherImpl{data: data}
}

func (p *PublisherImpl) checkTopic(topic string) bool {
	_, ok := p.data[topic]
	return ok
}

func (p *PublisherImpl) CreateTopic(topic string) error {
	if p.checkTopic(topic) {
		return ErrorTopicAlreadyExists
	} else {
		p.data[topic] = []message.Message{}
		return nil
	}
}

func (p *PublisherImpl) Publish(topic string, m message.Message) error {
	if p.checkTopic(topic) {
		p.data[topic] = append(p.data[topic], m)
		return nil
	} else {
		return ErrorTopicDoesNotExist
	}
}

func (p *PublisherImpl) Size(topic string) (int, error) {
	if v, ok := p.data[topic]; ok {
		return len(v), nil
	} else {
		return -1, ErrorTopicDoesNotExist
	}
}

func (p *PublisherImpl) Get(topic string, offset int) (message.Message, error) {
	if v, ok := p.data[topic]; ok {
		if len(v) > offset {
			return v[offset], nil
		} else {
			return message.Message{}, ErrorOffsetOverflow
		}
	} else {
		return message.Message{}, ErrorTopicDoesNotExist
	}
}
