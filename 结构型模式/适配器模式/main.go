package main

import "fmt"

// Adaptee is the interface that needs to be adapted
type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImpl struct{}

func (a *AdapteeImpl) SpecificRequest() string {
	return "specific request"
}

// Target is the target interface that we want to adapt to
type Target interface {
	Request() string
}

type Adapter struct {
	Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func Client(t Target) string {
	return t.Request()
}

func main() {
	adaptee := &AdapteeImpl{}
	target := &Adapter{Adaptee: adaptee}
	output := Client(target)
	fmt.Println(output)
}
