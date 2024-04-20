package main

import "fmt"

type Observer interface {
	Update(string)
}

type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(msg string) {
	fmt.Printf("%s received: %s\n", co.name, msg)
}

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Notify(string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (cs *ConcreteSubject) Register(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) Unregister(observer Observer) {
	for i, obs := range cs.observers {
		if obs == observer {
			cs.observers = append(cs.observers[:i], cs.observers...)
			break
		}
	}
}

func (cs *ConcreteSubject) Notify(msg string) {
	for _, observer := range cs.observers {
		observer.Update(msg)
	}
}

func main() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.Notify("Hello, Observers!")

	subject.Unregister(observer2)

	subject.Notify("Another message")
}
