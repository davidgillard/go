/* Moyenne et écart type de notes

a) Proposer une fonction moyenne qui admet comme argument d'entrée une liste non vide de notes. Cette fonction
retourne la moyenne des notes.

b) Proposer une fonction ecart_type qui admet comme argument d'entrée une liste non vide de notes. Cette
fonction retourne l'écart type des notes.

*/

package main

import (
	"fmt"
	"math"
)

func moyenne(notes []float64) float64 {

	somme := 0.0
	for _, v := range notes {
		somme += v
	}

	return somme / float64(len(notes))
}

func ecart_type(notes []float64) float64 {

	moy := moyenne(notes)
	var ecart_ty float64
	for j, _ := range notes {
		ecart_ty += math.Pow(notes[j]-moy, 2)
	}

	return math.Sqrt(ecart_ty / 10)
}

func main() {
	//notes := []float64{13, 13, 14, 12, 14, 12, 14, 15, 13, 12}
	notes := []float64{3, 13, 19, 18, 14, 2, 9, 15, 17, 1}
	fmt.Println("La moyenne des notes: ", moyenne(notes))
	fmt.Println("l'écart-type des notes: ", ecart_type(notes))
}
