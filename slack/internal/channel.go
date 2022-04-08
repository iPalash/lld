package internal

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
	users []User
	msgs  []Message
}

func NewChat(user1, user2 User) Chat {
	ch := &Channel{
		users: []User{user1, user2},
		msgs:  []Message{},
	}

	user1.Connect(ch, 0)
	user2.Connect(ch, 0)

	return ch
}

func (c *Channel) Send(msg Message) {
	c.msgs = append(c.msgs, msg)
	for _, u := range c.users {
		u.Notify()
	}
}

func (c *Channel) Size() int {
	return len(c.msgs)
}

func (c *Channel) Get(offset int) Message {
	return c.msgs[offset]
}

func NewGroup(user1 User) GroupChat {
	ch := &Channel{
		users: []User{user1},
		msgs:  []Message{},
	}

	user1.Connect(ch, 0)

	return ch
}

func (ch *Channel) Join(u User) {
	ch.users = append(ch.users, u)
	u.Connect(ch, ch.Size())
}

func (ch *Channel) Leave(u User) {
	for i, user := range ch.users {
		if u == user {
			ch.users = append(ch.users[:i], ch.users[i+1:]...)
		}
	}
	u.Disconnect(ch)
}
