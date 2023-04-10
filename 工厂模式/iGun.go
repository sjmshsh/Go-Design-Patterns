package main

// IGun 产品接口, 定义一支枪所需具备的所有方法
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
