package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	HandleRequest(string)
}

type ConcreteHandler1 struct {
	next Handler
}

func (h *ConcreteHandler1) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *ConcreteHandler1) HandleRequest(request string) {
	if request == "Request1" {
		fmt.Println("ConcreteHandler1 handled the request")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

type ConcreteHandler2 struct {
	next Handler
}

func (h *ConcreteHandler2) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *ConcreteHandler2) HandleRequest(request string) {
	if request == "Request2" {
		fmt.Println("ConcreteHandler2 handled the request")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func main() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}

	handler1.SetNext(handler2)

	handler1.HandleRequest("Request1")
	handler1.HandleRequest("Request2")
}
