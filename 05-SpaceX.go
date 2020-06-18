/*fichier : 05-SpaceX.go
auteur : David GILLARD */

/*

Le système de transport interplanétaire SpaceX n'a ​​pas d'entraînement de distorsion, mais il se rendra
sur Mars à un respectable 100 800 km / h. Une date de lancement ambitieuse de janvier 2025 placerait
Mars et la Terre distantes de 96 300 000 km. Combien de jours faudrait-il pour atteindre Mars?
Modifiez la liste 2.3 pour le savoir.

*/

package main // import du package main

import "fmt" // import de fmt du fournit les fonctions pour le formatage d'entrées / sorties

func main() { // le programme démarre avec la fonction du package main

	const nbresheuresparjours = 24
	var vitesse = 100800    // km/h
	var distance = 96300000 // km
	fmt.Println(distance/vitesse/nbresheuresparjours, "jours")
}
