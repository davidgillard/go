package main

import (
	"fmt"
)

func main() {

	age := 46
	ageSurMars := float64(age)

	nbreJoursSurMars := 687.0
	nbreJoursSurTerre := 365.2425
	ageSurMars = ageSurMars * nbreJoursSurTerre / nbreJoursSurMars
	fmt.Println("J'ai", ageSurMars, "ans sur la planéte Mars")
	fmt.Println("La convertion d'un float64 en entier donnr pour la variable nbreJoursSurTerre:", int(nbreJoursSurTerre))
}
