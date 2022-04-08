package internal

import (
	"testing"
	"time"
)

func TestNewChat(t *testing.T) {
	u1 := NewUser("1")
	u2 := NewUser("2")
	t.Run("base", func(t *testing.T) {
		chat := NewChat(u1, u2)
		chat.Send(u1.Message("hello"))
		time.Sleep(time.Millisecond * 10)
		chat.Send(u2.Message("world"))
	})
}

func TestNewGroup(t *testing.T) {
	u1 := NewUser("1")
	u2 := NewUser("2")
	u3 := NewUser("3")
	t.Run("base", func(t *testing.T) {
		group := NewGroup(u1)
		group.Send(u1.Message("hello"))
		time.Sleep(time.Millisecond * 10)
		group.Join(u2)
		group.Join(u3)
		time.Sleep(time.Millisecond * 10)
		group.Send(u2.Message("world"))
		group.Leave(u2)
		time.Sleep(time.Millisecond * 10)
		group.Send(u1.Message("hello again"))
	})
}