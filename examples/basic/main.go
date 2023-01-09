package main

import (
	"github.com/jmdavril/pubsub"
	"log"
)

////////////////////
// 1. Define event
////////////////////

const EMPLOYEE_CREATED_TYPE = "employee_created"

type EmployeeCreatedEvent struct {
	Firstname string
	Lastname  string
	Email     string
}

func (EmployeeCreatedEvent) Type() pubsub.EventType {
	return EMPLOYEE_CREATED_TYPE
}

//////////////////////
// 2. Define callback
//////////////////////

func StartSubscription(r pubsub.Sub) {
	eventHandler := func(e pubsub.Event) {
		pce, _ := e.(EmployeeCreatedEvent)
		log.Printf("Created employee '%s %s' with email '%s'", pce.Firstname, pce.Lastname, pce.Email)
		return
	}

	r.Subscribe(EMPLOYEE_CREATED_TYPE, eventHandler)
}

////////////////////
// 3. Publish event
////////////////////

func main() {
	p := pubsub.NewPublisher()
	StartSubscription(p)

	e := EmployeeCreatedEvent{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@email.com",
	}
	p.Publish(e)
}

