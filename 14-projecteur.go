/*fichier : 14-projecteur.go
auteur : David GILLARD */

/*
Le script suivant determine si 21OO est une année bissextile, deux rappels sont necessaires :
- une année est bissextile si elle est divisible par 4 et pas divisible par 100
- une année est bissextile si elle est divisible par 400
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
