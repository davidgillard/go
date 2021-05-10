package main

import (
	"fmt"
)

func main() {

	message := "shalom"
	c := message[4] // la caractére demandé sera la lettre o
	fmt.Printf("%c\n", c)

	// cette boucle affichera shalom avec un caractere par ligne
	for i := 0; i < 6; i++ {
		c = message[i]
		fmt.Printf("%c\n", c)
	}

}
