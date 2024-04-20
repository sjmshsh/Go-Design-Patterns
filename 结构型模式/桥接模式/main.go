package main

import "fmt"

// Shape 抽象部分 - 形状
type Shape interface {
	Draw()
}

// Color 实现部分 - 颜色
type Color interface {
	ApplyColor() string
}

type RedColor struct{}

func (r *RedColor) ApplyColor() string {
	return "Red"
}

type BlueColor struct{}

func (b *BlueColor) ApplyColor() string {
	return "Blue"
}

// Circle 桥接部分 - 形状和颜色的组合
type Circle struct {
	color Color
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a %s circle\n", c.color.ApplyColor())
}

type Square struct {
	color Color
}

func (s *Square) Draw() {
	fmt.Printf("Drawing a %s square\n", s.color.ApplyColor())
}

func main() {
	red := &RedColor{}
	blue := &BlueColor{}

	circleRed := &Circle{color: red}
	circleBlue := &Circle{color: blue}

	squareRed := &Square{color: red}
	squareBlue := &Square{color: blue}

	circleRed.Draw()
	circleBlue.Draw()
	squareRed.Draw()
	squareBlue.Draw()
}
