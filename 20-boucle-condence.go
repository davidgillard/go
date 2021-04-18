/*
Une façon condencé d'écrire une boucle for
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	for count := 20; count > 10; count-- {
		time.Sleep(time.Second)
		fmt.Println(count)
	}
}
