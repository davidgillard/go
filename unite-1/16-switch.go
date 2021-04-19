/*fichier : 16-switch.go
auteur : David GILLARD */

/*
Go dispose d'un outil permettant de comparer une valeur à plusieurs
autres à travers le traitement "switch" */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("l'entrée de la caverne est ici ")
	var action = "entrez dans la caverne"

	switch action {
	case "sors":
		fmt.Println("je commence à flipper, je préfére rebrousser chemin")
	case "entrez dans la caverne":
		fmt.Println("Oh, je suis un homme quand même !")
	case "lire les panneaux":
		fmt.Println("Seuls les adultes conciens sont censé pénétres les lieux")
	}
}
