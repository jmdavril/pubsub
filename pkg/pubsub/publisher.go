package pubsub

import (
	"sync"
)

type EventType string

type Subscription func(Event)

// Register is used by a subscriber to subscribe to a particular EventType and register a callback of type Subscription
type Register interface {
	Subscribe(et EventType, s Subscription)
}

// Producer is used to push events to subscribers who have subscribed to e.Type
type Producer interface {
	Push(e Event)
}

// Publisher implements interfaces Register and Producer
type Publisher struct {
	mu            sync.RWMutex
	subscriptions map[EventType][]Subscription
}

// Verify interfaces compliance at compile time
var _ Register = (*Publisher)(nil)
var _ Producer = (*Publisher)(nil)

func NewPublisher() *Publisher {
	return &Publisher{
		subscriptions: make(map[EventType][]Subscription),
	}
}

func (p *Publisher) Subscribe(et EventType, s Subscription) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscriptions[et] = append(p.subscriptions[et], s)
}

func (p *Publisher) Push(e Event) {
	p.mu.Lock()
	subs := p.subscriptions[e.Type()]
	p.mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(len(subs))

	for _, sub := range subs {
		go func(s Subscription) {
			defer wg.Done()
			s(e)
		}(sub)
	}

	wg.Wait()
}
