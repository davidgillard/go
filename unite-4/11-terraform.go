package main

import (
	"fmt"
)

type Planete []string

func (planete Planete) terraform() {
	for i := range planete {
		planete[i] = "New" + planete[i]
	}
}

func main() {
	planete := []string{
		"Mercure", "Venus", "Terre", "Mars",
		"Neptune", "Uranus", "Saturne", "Jupiter",
	}
	Planete(planete[3:4]).terraform()
	Planete(planete[6:]).terraform()
	fmt.Println(planete)
}
