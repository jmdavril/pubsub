package pubsub

type Event interface {
	Type() EventType
}
