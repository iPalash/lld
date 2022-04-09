package internal

import (
	"sync"
)

type Chat interface {
	Send(Message)
	Size() int
	Get(int) Message
}

type GroupChat interface {
	Send(Message)
	Size() int
	Get(int) Message
	Join(User)
	Leave(User)
}

type Channel struct {
	users    []User
	msgs     []Message
	add      chan User
	leave    chan User
	incoming chan Message
	stop     chan struct{}
	lock     *sync.RWMutex
}

func NewChat(user1, user2 User) Chat {
	ch := NewGroup(user1)

	ch.Join(user2)

	return ch
}

func (c *Channel) Send(msg Message) {
	c.incoming <- msg
}

func (c *Channel) _send(msg Message) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.msgs = append(c.msgs, msg)
	for _, u := range c.users {
		go func(u User) {
			u.Notify(c)
		}(u)
	}
}

func (c *Channel) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.msgs)
}

func (c *Channel) Get(offset int) Message {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.msgs[offset]
}

func NewGroup(user1 User) GroupChat {
	ch := &Channel{
		users:    []User{user1},
		msgs:     []Message{},
		add:      make(chan User),
		leave:    make(chan User),
		incoming: make(chan Message),
		stop:     make(chan struct{}),
		lock:     &sync.RWMutex{},
	}

	user1.Connect(ch, 0)
	go func() {
		ch.handleEvents()
	}()

	return ch
}

func (ch *Channel) Join(u User) {
	ch.add <- u
}

func (ch *Channel) _join(u User) {

	ch.users = append(ch.users, u)
	u.Connect(ch, ch.Size())

}
func (ch *Channel) Leave(u User) {
	ch.leave <- u
}

func (ch *Channel) _leave(u User) {
	for i, user := range ch.users {
		if u == user {
			ch.users = append(ch.users[:i], ch.users[i+1:]...)
		}
	}
	u.Disconnect(ch)
	if len(ch.users) == 0 {
		ch.stop <- struct{}{}
	}
}

func (ch *Channel) handleEvents() {
	loop := true
	for loop {
		select {
		case u := <-ch.add:
			ch._join(u)
		case u := <-ch.leave:
			ch._leave(u)
		case msg := <-ch.incoming:
			ch._send(msg)
		case <-ch.stop:
			loop = false
		}
	}

}
