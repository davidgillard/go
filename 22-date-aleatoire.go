/*
le script suivant génére une année aléatoire au lieu d'utiliser une constante
pour le mois de février on considére qu'il est de 29 jours pour les années bissextiles et de 28 pour les autres
Note :
- Vous pouvez mettre un "if" à l'interieur du bloc "case"
- Utiliser une boucle for pour afficher 10 dates aléatoires
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {

	annee := rand.Intn(10) + 2021
	mois := rand.Intn(12) + 1
	nbreJourDansMois := 31

	switch mois { // ici les variables era - annee - mois sont accessibles
	case 2:
		if annee%400 == 0 || annee%4 == 0 && annee%100 != 0 {
			nbreJourDansMois = 29
		} else {
			nbreJourDansMois = 28
		}

	case 4, 6, 9, 11:
		nbreJourDansMois = 30

	}

	jour := rand.Intn(nbreJourDansMois) + 1
	for count := 0; count < 10; count++ {
		time.Sleep(time.Second)
		fmt.Println(era, annee, mois, jour)
	}
}
