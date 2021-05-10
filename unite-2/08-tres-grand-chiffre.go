package main

import (
	"fmt"
	"math/big"
)

func main() {
	vitesseLumiere := big.NewInt(299792)
	joursEnSecondes := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("La galaxie d'androméde est à ", distance, "kms de la Terre")

	secondes := new(big.Int)
	secondes.Div(distance, vitesseLumiere)

	jours := new(big.Int)
	jours.Div(secondes, joursEnSecondes)
	//const joursParAn = 365

	//annees := jours / joursParAn
	fmt.Println("C'est", jours, "jours de voyage à la vitesse de la lumière.")
	//fmt.Println("soit", annees, "années à la vitesse de la lumière.")
}
