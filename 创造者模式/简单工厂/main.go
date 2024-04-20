package main

import "fmt"

type Product interface {
	GetName() string
}

type ProductA struct{}

func (p *ProductA) GetName() string {
	return "Product A"
}

type ProductB struct{}

func (p *ProductB) GetName() string {
	return "Product B"
}

type Factory struct{}

func (f *Factory) CreateProduct(name string) Product {
	switch name {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		return nil
	}
}

func main() {
	factory := &Factory{}
	productA := factory.CreateProduct("A")
	productB := factory.CreateProduct("B")
	fmt.Println(productA.GetName())
	fmt.Println(productB.GetName())
}
