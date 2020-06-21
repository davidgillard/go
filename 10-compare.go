/*fichier : 10-compare.go
auteur : David GILLARD */

/*
opérateur de comparaison
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Il y a un panneau informatif qui mentionne que l'entrée est interdite aux mineurs.")
	var (
		age    = 41
		mineur = 18 < age
	)
	fmt.Printf("à l'age de %v ans, suis-je un mineur ? %v\n", age, mineur)
}
