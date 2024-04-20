package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) SetName(name string) *Person {
	p.name = name
	return p
}

func (p *Person) SetAge(age int) *Person {
	p.age = age
	return p
}

func (p *Person) Describe() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func main() {
	person := NewPerson("John", 30)
	person.SetName("Jack").SetAge(25).Describe()
}
