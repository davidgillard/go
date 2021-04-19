/*fichier : 07-distance_aleatoire.go
auteur : David GILLARD */

/*
La distance existante entre la Terre et Mars du soleil. Écrire un programme qui génère une distance
aléatoire de 56 000 000 à 401 000 000 km.
*/

package main // import du package main

import (
	"fmt" // import de fmt du fournit les fonctions pour le formatage d'entrées / sorties
	"math/rand"
)

func main() { // le programme démarre avec la fonction du package main

	var distance = rand.Intn(345000001) + 56000000
	fmt.Println(distance)

}
