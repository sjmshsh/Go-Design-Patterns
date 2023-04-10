package main

// Gun 实现了iGun接口的结构体类型
type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}

func (g *Gun) GetPower() int {
	return g.Power
}
