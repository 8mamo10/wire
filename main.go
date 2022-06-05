package main

import "fmt"

type Message string

type Greeter struct {
	message Message
}

type Event struct {
	greeter Greeter
}

func NewMessage() Message {
	return Message("Hi there")
}

func NewGreeter(m Message) Greeter {
	return Greeter{message: m}
}

func (g Greeter) Greet() Message {
	return g.message
}

func NewEvent(g Greeter) Event {
	return Event{greeter: g}
}

func (e Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()
}
