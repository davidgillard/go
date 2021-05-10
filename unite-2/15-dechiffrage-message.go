/*
Decipher the quote from Julius Caesar:
L fdph, L vdz, L frqtxhuhg.
—Julius Caesar
Your program will need to shift uppercase and lowercase letters by –3. Remember that
'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z', and likewise for uppercase letters.
*/

package main

import (
	"fmt"
)

func main() {

	message := "L fdph, L vdz, L frqtxhuhg."
	fmt.Println("Le message original :", message)
	fmt.Println("Le message est composé de ", len(message), "caratères")
	fmt.Println("!!! Dé-chiffrement du message !!!")
	for i := 0; i < len(message); i++ {
		c := message[i]
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
