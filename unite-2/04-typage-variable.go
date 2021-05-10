package main

import (
	"fmt"
)

func main() {

	a := "Procrastination"
	fmt.Printf("Le type de la variable est: %T pour le contenu suivant: %[1]v\n", a)

	b := 75
	fmt.Printf("Le type de la variable est: %T pour le contenu suivant: %[1]v\n", b)

	c := 75.5
	fmt.Printf("Le type de la variable est: %T pour le contenu suivant: %[1]v\n", c)

	d := true
	fmt.Printf("Le type de la variable est: %T pour le contenu suivant: %[1]v\n", d)

	var rouge, vert, bleu uint8 = 0, 141, 213

	fmt.Printf("%x %x %x\n", rouge, vert, bleu)

	fmt.Printf("color: #%02x%02x%02x;\n", rouge, vert, bleu)
}
