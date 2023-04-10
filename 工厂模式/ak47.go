package main

// Ak47 具体的枪支
type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "Ak47 gun",
			Power: 4,
		}}
}
