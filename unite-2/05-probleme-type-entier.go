/*
Cet exercice montre les limites du typage des nombres entiers
c'est la raison pour laquelle il ne faut pas hésiter à passer au typage
supérieur si on se rencontre de certaines limites
*/

package main

import (
	"fmt"
)

func main() {
	var blue uint8 = 255       // la valeur bleu est definit à 255 sur un type unit8 donc de 0 à 255
	fmt.Printf("%08b\n", blue) // en binaire cela donne 11111111
	blue++                     // ici on incrémente 1 unité du coup que se passe t'il ?
	fmt.Printf("%08b\n", blue) // le resultat est le suivant 00000000
}
