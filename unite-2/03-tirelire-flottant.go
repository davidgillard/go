/*
Économisez de l'argent pour acheter un cadeau pour votre ami. Écrivez un programme qui place au hasard
Cinq centimes (0,05€), dix centimes (0,10€) et vingt cinq cents (0,50 €) dans une tirelire vide jusqu'à ce qu'il
contiennent au moins 20,00€. Afficher le solde courant de la tirelire après chaque dépôt,
le formater avec une largeur et une précision appropriées.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tirelire := 0.0

	for tirelire < 20 {
		seconds := time.Now().Unix()
		rand.Seed(seconds)
		switch rand.Intn(3) {
		case 0:
			tirelire += 0.05
		case 1:
			tirelire += 0.10
		case 2:
			tirelire += 0.50
		}
		fmt.Printf("%5.2f\n", tirelire)
	}
}
