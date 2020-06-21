/*fichier : 12-jeu.go
auteur : David GILLARD */

/*
Les jeux d'aventure sont divisés en salles. Écrivez un programme qui utilise la structure if
et else if pour afficher une description de chacune des trois pièces: grotte, entrée et montagne.
Lors de l'écriture de votre programme, assurez-vous que les accolades {} soient placées en fonction
de celle vraie style de renfort
*/

package main

import (
	"fmt"
)

func main() {
	var salle = "pole nord"
	if salle == "grotte" {
		fmt.Println("Vous êtes dans une grotte")
	} else if salle == "entrée" {
		fmt.Println("Vous êtes à l'entrée de la cave")
	} else if salle == "montagne" {
		fmt.Println("Il y a une falaise ici. Un chemin mène vers l'ouest en bas de la montagne.")
	} else {
		fmt.Println("Tout est blanc")
	}
}
