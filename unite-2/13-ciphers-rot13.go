package main

import (
	"fmt"
)

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	fmt.Println(len(message))
	fmt.Println("!!! Dé-chiffrement du message !!!")

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}

		fmt.Printf("%c", c)
	}
}
