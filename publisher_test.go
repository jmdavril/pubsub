package pubsub

import "testing"

const (
	type1 = "Type 1"
	type2 = "Type 2"
)

type event1 struct {
	value int
}

func (e event1) Type() EventType {
	return type1
}

type event2 struct {
	value int
}

func (e event2) Type() EventType {
	return type2
}

func TestPublisher(t *testing.T) {

	// Arrange
	var a int
	var b int

	fa := func(e Event) {
		e1, _ := e.(event1)
		a += e1.value
	}

	fb := func(e Event) {
		e1, _ := e.(event1)
		b += e1.value
	}

	fc := func(e Event) {
		e2, _ := e.(event2)
		a += e2.value
	}

	e1a := event1{
		value: 1,
	}

	e1b := event1{
		value: 5,
	}

	e2 := event2{
		value: 10,
	}

	p := NewPublisher()

	// Act
	p.Subscribe(type1, fa)
	p.Subscribe(type1, fb)
	p.Subscribe(type2, fc)

	p.Publish(e1a)
	p.Publish(e1b)
	p.Publish(e2)

	// Assert
	if a != 16 {
		t.Fatalf("a is wrong. expected=%d, got=%d", 16, a)
	}
	if b != 6 {
		t.Fatalf("a is wrong. expected=%d, got=%d", 6, b)
	}
}
