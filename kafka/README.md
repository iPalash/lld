# Problem Statement
- Implement a Multithreaded queue like with multiple publishers and subscriber over same topic

## Requirements
- Should support multiple topics 
- Publishers should be able to publish to a topic
- Multiple subcribers can consume from the same topic without blocking one another
![Flow](flow.png)

## Interfaces
```go
type Publisher interface {
	CreateTopic(string) (Topic, error)
	Publish(string, Message) error
}

type Subscriber interface {
	Start()
	Notify()
}

type Topic interface {
	Add(Message)
	AddSubscription(Subscriber)
	Size() int
	Get(int) Message
}
```

