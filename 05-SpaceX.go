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

	const vitessedelalumiere = 299792 // affectation de 299792 a la constante vitessedelalumiere en km/s
	var distance = 56000000           // affectation de 56000000 à la variable distance en km
	fmt.Println(distance/vitessedelalumiere, "secondes")
	distance = 401000000
	fmt.Println(distance/vitessedelalumiere, "seconds")
}
