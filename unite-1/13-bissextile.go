/*fichier : 13-bissextile.go
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

	fmt.Println("L'année 2100 est elle une année bissextile ? ")
	var (
		annee      = 2100
		bissextile = annee%400 == 0 || (annee%4 == 0 && annee%100 != 0)
	)

	if bissextile {
		fmt.Printf("%v est une année bissextile\n", annee)
	} else {
		fmt.Printf("%v n'est pas une année bissextile\n", annee)
	}
}
