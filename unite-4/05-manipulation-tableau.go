package main

import "fmt"

// terraform accomplishes nothing
func terraform(planetes [8]string) {
	for i := range planetes {
		planetes[i] = "New " + planetes[i]
	}
}

func main() {
	planetes := [...]string{
		"Jupiter",
		"Mars",
		"Mercure",
		"Neptune",
		"Saturne",
		"Terre",
		"Uranus",
		"Venus",
	}

	terraform(planetes)
	fmt.Println(planetes)

}
