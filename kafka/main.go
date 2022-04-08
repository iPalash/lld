package main

import (
	"fmt"
	"kafka/queue"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := queue.New()
	topicName := "topic1"
	topic, _ := p.CreateTopic(topicName)

	queue.NewSub(topic).Start()

	queue.NewSub(topic).Start()

	//Give some time for subs to start
	time.Sleep(time.Millisecond * 10)

	for i := 0; i < 5; i++ {
		go func(idx int) {
			p.Publish(topicName, queue.NewMessage(fmt.Sprint(idx)))
		}(i)
	}

	p.Publish(topicName, queue.NewMessage("world"))
	time.Sleep(time.Second)

	queue.NewSub(topic).Start()

	time.Sleep(time.Second * 10)
}
