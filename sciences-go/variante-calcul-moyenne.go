/* Moyenne et écart type de notes

a) Proposer une fonction moyenne qui admet comme argument d'entrée une liste non vide de notes. Cette fonction
retourne la moyenne des notes.

b) Proposer une fonction ecart_type qui admet comme argument d'entrée une liste non vide de notes. Cette
fonction retourne l'écart type des notes.

*/

package main

import "fmt"

func moyenne(notes []float64) float64 {

	somme := 0.0
	for elt := 0; elt < len(notes); elt++ {
		somme += notes[elt]

	}
	return somme / float64(len(notes))
}

func main() {
	notes := []float64{2, 13, 9, 17, 6, 15}
	fmt.Println(moyenne(notes))
}
