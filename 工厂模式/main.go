package main

import "fmt"

func main() {
	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.GetName())
	fmt.Printf("Power: %d\n", g.GetPower())
}

// Gun: Ak47 gun
// Power: 4
// Gun: Musket gun
// Power: 1
