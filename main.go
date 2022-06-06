package main

import (
	"fmt"
	"os"
	"time"
)

type Message string

type Greeter struct {
	message Message
	grumpy  bool
}

type Event struct {
	greeter Greeter
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{message: m, grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.grumpy {
		return Message("Go away!")
	}
	return g.message
}

func NewEvent(g Greeter) (Event, error) {
	if g.grumpy {
		return Event{}, fmt.Errorf("could not create event: event greeter is grumpy")
	}
	return Event{greeter: g}, nil
}

func (e Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent("Come here")
	if err != nil {
		fmt.Printf("failed to create event: %v\n", err)
		os.Exit(2)
	}
	e.Start()
}
