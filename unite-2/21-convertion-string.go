package main

import (
	"fmt"
)

func main() {
	yesNo := "yes"
	var lancer bool

	switch yesNo {
	case "true", "yes", "1":
		lancer = true
	case "false", "no", "0":
		lancer = false
	default:
		fmt.Println(yesNo, "n'est pas valide")

	}
	fmt.Println("Prêt pour le lancement :", lancer)
}
