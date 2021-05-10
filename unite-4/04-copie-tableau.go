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

	fmt.Println("Valeur d'origine du tableau planete: ", planete)

	uneCopiedePlanete := planete // uneCopiedePlanete est ici une copie du tableau planete
	planete[2] = "Exoplanète"    // on affecte au deuxième élément du tableau planete la valeur "Exoplanète" en lieu et place de "Mercure"

	fmt.Println("Affichage de planete apres modification: ", planete)
	fmt.Println("################################################################################")
	fmt.Println("Affichage de uneCopiedePlanete: ", uneCopiedePlanete)
}
