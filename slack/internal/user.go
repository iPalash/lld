package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User interface {
	Connect(Chat, int)
	Notify(Chat)
	Message(string) Message
	Disconnect(Chat)
}

type UserImpl struct {
	Name       string
	ID         int
	offset     int
	block      *sync.Map
	disconnect *sync.Map
}

func NewUser(name string) User {
	return &UserImpl{
		Name:       name,
		ID:         rand.Int(),
		offset:     0,
		block:      &sync.Map{},
		disconnect: &sync.Map{},
	}
}

func (u *UserImpl) Connect(ch Chat, offset int) {
	stop := make(chan struct{})
	u.disconnect.Store(ch, stop)

	unblock := make(chan struct{})
	u.block.Store(ch, unblock)
	go func(offset int, chn chan struct{}, unblock chan struct{}) {
		u.offset = offset
		loop := true
		for loop {
			for u.offset < ch.Size() {
				u.receiveMessage(ch.Get(u.offset))
				u.offset += 1
			}
			select {
			case <-unblock:
				continue
			case <-stop:
				loop = false
			}

		}
	}(offset, stop, unblock)
}

func (u *UserImpl) Disconnect(ch Chat) {
	stop, _ := u.disconnect.LoadAndDelete(ch)
	stop.(chan struct{}) <- struct{}{}
}

func (u *UserImpl) receiveMessage(msg Message) {
	if msg.user != u.ID {
		fmt.Printf("user:%s got messsage %v\n", u.Name, msg)
	}
}

func (u *UserImpl) Notify(ch Chat) {
	unblock, _ := u.block.Load(ch)
	unblock.(chan struct{}) <- struct{}{}
}

func (u *UserImpl) Message(text string) Message {
	return Message{
		text: text,
		user: u.ID,
		time: time.Now().UnixNano(),
	}
}
