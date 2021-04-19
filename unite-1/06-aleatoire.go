/*fichier : 06-aleatoire.go
auteur : David GILLARD */

/*
Demandez maintenant à votre ordinateur de «penser» à un nombre compris entre 1 et 10. Votre ordinateur
peut générer des nombres pseudo-aléatoires à l'aide du package rand. Ils sont appelés pseudo-aléatoires
car ils sont plus ou moins aléatoires, mais pas vraiment aléatoires.
Le code de la liste 2.6 affichera deux nombres entre 1 et 10. Passer 10 à Intn revient
un nombre compris entre 0 et 9, auquel vous ajoutez 1 et affectez le résultat à num. La variable num ne
peut pas être une constante Go parce que c'est le résultat d'un appel de fonction.
*/

package main // import du package main

import (
	"fmt" // import de fmt du fournit les fonctions pour le formatage d'entrées / sorties
	"math/rand"
	"time"
)

func main() { // le programme démarre avec la fonction du package main
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	num := rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)
}
