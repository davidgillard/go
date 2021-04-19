/*
le script suivant mettra en lumière l'importance de la portée des variables
*/

package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {

	annee := 2021
	mois := rand.Intn(12) + 1
	nbreJourDansMois := 31

	switch mois { // ici les variables era - annee - mois sont accessibles
	case 2:
		nbreJourDansMois = 28
	case 4, 6, 9, 11:
		nbreJourDansMois = 30

	}

	jour := rand.Intn(nbreJourDansMois) + 1
	fmt.Println(era, annee, mois, jour)
}
