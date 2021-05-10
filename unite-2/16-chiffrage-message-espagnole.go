/*
Cipher the Spanish message “Hola Estación Espacial Internacional” with ROT13. Mod-
ify listing 9.7 to use the range keyword. Now when you use ROT13 on Spanish text, char-
acters with accents are preserved.
*/

package main

import (
	"fmt"
)

func main() {

	message := "Hola Estación Espacial Internacional"
	for _, c := range message {

		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
}
