package main

import (
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
	subscriber.New(p, topic).Start()

	p.Publish(topic, message.New("hello"))
	p.Publish(topic, message.New("world"))
	time.Sleep(time.Minute)
}
