package main

import (
	"fmt"
)

func main() {
	planetes := [...]string{
		"Mercure",
		"Venus",
		"Terre",
		"Mars",
		"Jupiter",
		"Saturne",
		"Uranus",
		"Neptune",
	}

	telluriques := planetes[0:4]
	gazeuses := planetes[4:6]
	glaciers := planetes[6:8]

	fmt.Println(telluriques, gazeuses, glaciers)
	fmt.Println(glaciers[0])
}
