package main

import (
	"fmt"
	"kafka/message"
	"kafka/publisher"
	"kafka/subscriber"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := publisher.New()
	topic := "t1"
	p.CreateTopic(topic)
	// p.CreateTopic("t1")
	subscriber.New(p, topic).Start()
	subscriber.New(p, topic).Start()

	for i := 0; i < 5; i++ {
		go func(idx int) {
			p.Publish(topic, message.New(fmt.Sprint(i)))

		}(i)
	}

	p.Publish(topic, message.New("world"))
	time.Sleep(time.Second * 10)
}
