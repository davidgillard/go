/*fichier : 11-branches.go
auteur : David GILLARD */

/*
opérateur de comparaison
*/

package main

import (
	"fmt"
)

func main() {
	var command = "Aller à l'est"
	if command == "Aller à l'est" {
		fmt.Println("Vous vous dirigez plus haut sur la montagne.")
	} else if command == "aller à l'intérieur" {
		fmt.Println("Vous entrez dans la grotte où vous vivez le reste de votre vie.")
	} else {
		fmt.Println("Je n'ai pas vraiment compris ça.")
	}

}
