package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
	age  int
}

func (cp *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		name: cp.name,
		age:  cp.age,
	}
}

func NewPrototype(name string, age int) Prototype {
	return &ConcretePrototype{
		name: name,
		age:  age,
	}
}

func main() {
	prototype := NewPrototype("Bob", 30)
	clone1 := prototype.Clone()
	clone2 := prototype.Clone()

	fmt.Println(prototype)
	fmt.Println(clone1)
	fmt.Println(clone2)
}
