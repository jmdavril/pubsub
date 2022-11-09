package main

import (
	"github.com/jmdavril/pubsub/pkg/pubsub"
	"log"
)

// service 1 pushes an event e with e.value = "hello", service 2 is the subscriber and prints e.value when receiving e
func main() {
	p := pubsub.NewPublisher()

	s1 := &service1{
		producer: p,
	}

	s2 := &service2{
		register: p,
	}

	s2.startSubscription()
	s1.pushValueHello()
}

const eventType = "example.basic"

type basicEvent struct {
	value string
}

func (e basicEvent) Type() pubsub.EventType {
	return eventType
}

type service1 struct {
	producer pubsub.Producer
}

func (s service1) pushValueHello() {
	s.producer.Push(basicEvent{value: "hello"})
}

type service2 struct {
	register pubsub.Register
}

func (s service2) startSubscription() {
	f := func(e pubsub.Event) {
		be := e.(basicEvent)
		log.Printf("Received message '%s'", be.value)
	}
	s.register.Subscribe(eventType, f)
}

