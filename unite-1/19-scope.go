/*
exercice concernant la portée des variables
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var count = 0
	var num = 0
	for count < 10 {
		num = rand.Intn(10) + 1
		fmt.Println(num)
		count += 1
	}
	fmt.Println("La valeur de num en dehors de la boucle: ", num)
}
