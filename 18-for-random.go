/*fichier : 19-for-random.go
auteur : David GILLARD */

/*
décollage
Écrivez un programme qui devine un nombre. Faites en sorte que l'ordinateur choisisse des nombres aléatoires entre 1 et 100
jusqu'à ce qu'il devine votre nombre, que vous déclarez en haut du programme.
Affichez chaque estimation et si elle était trop grande ou trop petite.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var number = 46
	for number > 0 {
		fmt.Println(number)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		number -= 1
	}
	if number == 0 {
		fmt.Println("!!! Décollage !!!")
	} else {
		fmt.Println("!!! Lancement de la fusée échouée !!!")
	}

}
