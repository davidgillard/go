package main

import (
	"fmt"
)

func main() {

	planete := [...]string{
		"Jupiter",
		"Mars",
		"Mercure",
		"Neptune",
		"Saturne",
		"Terre",
		"Uranus",
		"Venus",
	}

	fmt.Println(planete)

	// ##### solution n°1 on utilise for d'une certaine façon qui fournit plus de fléxibilité #####
	// for i := 0; i < len(planete); i++ {
	// 	laVoixLactee := planete[i]
	// 	fmt.Println(i, laVoixLactee)
	//}

	// #### solution n°2 on utilise un for avec l'option range qui simplifie la boucle mais donne moins de possibilité
	for i, planete := range planete {
		fmt.Println(i, planete)
	}
	// les deux solutions apportent le même rendu

}
