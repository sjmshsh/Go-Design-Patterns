package main

import "fmt"

// Component 是一个接口, 定义了一个操作
type Component interface {
	Operation() string
}

// ConcreteComponent 是一个具体的组件, 实现了Component接口
type ConcreteComponent struct{}

func (C *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator 是一个装饰器函数, 接收一个Component接口, 并返回一个新的带有附加行为的组件
type Decorator func(component Component) Component

// ConcreteDecoratorA 是一个具体的装饰器, 给Component添加新的行为
func ConcreteDecoratorA(component Component) Component {
	return &decoratorA{component: component}
}

type decoratorA struct {
	component Component
}

func (d *decoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

func ConcreteDecoratorB(component Component) Component {
	return &decoratorB{component: component}
}

type decoratorB struct {
	component Component
}

func (d *decoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}

func main() {
	component := &ConcreteComponent{}

	decoratedComponentA := ConcreteDecoratorA(component)
	decoratedComponentB := ConcreteDecoratorB(decoratedComponentA)

	fmt.Println(component.Operation())
	fmt.Println(decoratedComponentA.Operation())
	fmt.Println(decoratedComponentB.Operation())
}
