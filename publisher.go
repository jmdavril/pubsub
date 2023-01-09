package pubsub

import (
	"sync"
)

type EventType string

type Event interface {
	Type() EventType
}

type Subscription func(Event)

// Sub is used by a subscriber to subscribe to a particular EventType and register a callback of type Subscription
type Sub interface {
	Subscribe(et EventType, s Subscription)
}

// Pub is used to publish events to subscribers who have subscribed to e.Type
type Pub interface {
	Publish(e Event)
}

// PubSub implements interfaces Sub and Pub
type PubSub struct {
	mu            sync.RWMutex
	subscriptions map[EventType][]Subscription
}

// Verify interfaces compliance at compile time
var _ Sub = (*PubSub)(nil)
var _ Pub = (*PubSub)(nil)

func NewPublisher() *PubSub {
	return &PubSub{
		subscriptions: make(map[EventType][]Subscription),
	}
}

func (p *PubSub) Subscribe(et EventType, s Subscription) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscriptions[et] = append(p.subscriptions[et], s)
}

func (p *PubSub) Publish(e Event) {
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
