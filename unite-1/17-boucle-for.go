/*fichier : 16-boucle-for.go
auteur : David GILLARD */

/*
Plutôt que d'avoir à réécrire chacune des actions que l'on veux répété voyons ce que nous apporte la
boucle for : Exemple de traitement avec la bouche for  */

package main

import (
	"fmt"
	"time"
)

func main() {

	var count = 10
	for count > 0 {
		fmt.Println("Voici la valeur de la variable count: ", count)
		time.Sleep(time.Second)
		count -= 1
	}
	fmt.Println("!!! Décollage de la fusée !!!")
}
