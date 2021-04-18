/*fichier : 14-projecteur.go
auteur : David GILLARD */

/*
Le script suivant va concerné les booléens
*/

package main

import (
	"fmt"
)

func main() {

	var (
		avoirProjecteur  = true
		projecteurAllume = false
	)

	if !avoirProjecteur || !projecteurAllume {
		fmt.Println("Il n'est pas possible de voir quelque chose ici")
	}
}
