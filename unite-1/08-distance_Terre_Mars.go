/*fichier : 08-distance_Terre_Mars.go
auteur : David GILLARD */

/*

Écrire un programme qui détermine à quelle vittese (km/h) à besoin de voyager une navette
pour se rendre sur Mars depuis la Terre en 28 jours. Sachant que la distance est de 56 000 000 km.

*/

package main // import du package main

import (
	"fmt" // import de fmt du fournit les fonctions pour le formatage d'entrées / sorties
)

func main() { // le programme démarre avec la fonction du package main

	const nbresheuresparjours = 24
	var (
		jours    = 28       // km/h
		distance = 56000000 // km
	)
	fmt.Println(distance/(jours*nbresheuresparjours), "km/h")

}
