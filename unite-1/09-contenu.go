/*fichier : 08-contenu.go
auteur : David GILLARD */

/*
Ce programme verifie si une chaine de caractére contient un mot en particulier dans le code ci dessous
si la variable chaine contient le mot exterieur
*/

package main // import du package main

import (
	"fmt" // import de fmt du fournit les fonctions pour le formatage d'entrées / sorties
	"strings"
)

func main() { // le programme démarre avec la fonction du package main

	fmt.Println("Vous êtes dans une grotte.")
	var (
		chaine = "Vous marchez à l'exterieur de la grotte"
		sortie = strings.Contains(chaine, "exterieur")
		//portdelunettedesoleil = true
		//information           = strings.Contains(chaine, "information")
	)

	fmt.Println("Vous quittez la grotte.", sortie)

}
