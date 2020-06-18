/*fichier : vitessedelalumiere.go
auteur : David GILLARD */

/*
ce programme va calculer la durée du voyage entre la Terre et Mars

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
