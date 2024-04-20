package main

import "fmt"

type Product struct {
	PartA string
	PartB string
	PartC string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// ConcreteBuilder 具体建造者类
type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.PartA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.PartB = "Part B"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.PartC = "Part C"
}

func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	concreteBuilder := &ConcreteBuilder{product: &Product{}}
	director := &Director{builder: concreteBuilder}
	director.Construct()
	product := concreteBuilder.GetProduct()
	fmt.Println(product.PartA, product.PartB, product.PartC)
}
