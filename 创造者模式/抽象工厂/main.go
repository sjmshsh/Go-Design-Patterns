package main

import "fmt"

type Button interface {
	Click()
}

type Checkbox interface {
	Check()
}

// GUIFactory 抽象工厂
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// WinButton 具体产品
type WinButton struct{}

func (w *WinButton) Click() {
	fmt.Println("Windows button clicked")
}

type WinCheckbox struct{}

func (w *WinCheckbox) Check() {
	fmt.Println("Windows checkbox checked")
}

type MacButton struct{}

func (m *MacButton) Click() {
	fmt.Println("Mac button clicked")
}

type MacCheckbox struct{}

func (m *MacCheckbox) Check() {
	fmt.Println("Mac checkbox checked")
}

// WinFactory 具体工厂
type WinFactory struct{}

func (w *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (w *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (m *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

func main() {
	var factory GUIFactory

	// 根据需要选择工厂
	factory = &WinFactory{}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Click()
	checkbox.Check()
}
