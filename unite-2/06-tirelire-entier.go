/*
Écrivez un nouveau programme de tirelire qui utilise des entiers pour suivre le nombre de cents plutôt
que de €uros. Placez au hasard des pièces de 5 cents, de dix cents et des pièces de 50 cents dans une tirelire
vide jusqu'à ce qu'elle contienne au moins 20€. Affichez le solde courant de la tirelire après chaque dépôt en €uros
(par exemple, 1,05€).

CONSEIL Si vous avez besoin de trouver le reste de la division de deux nombres, utilisez le module (%).
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	tirelire := 0

	for tirelire < 2000 {
		seconds := time.Now().Unix()
		rand.Seed(seconds)
		switch rand.Intn(3) {
		case 0:
			tirelire += 5
		case 1:
			tirelire += 10
		case 2:
			tirelire += 20
		}
		euros := tirelire / 100
		cents := tirelire % 100
		fmt.Printf("%d.%0.2d€\n", euros, cents)
	}
}
