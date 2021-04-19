/*
Bienvenue au premier défi. Il est temps de prendre tout ce qui est abordé dans l’unité 1 et d’écrire
vous-même un programme. Votre défi est d'écrire un générateur de tickets dans le Go Playground qui utilise
des variables, des constantes. Il doit également s'appuyer sur les packages fmt et math/rand pour afficher et
aligner du texte et générer des nombres aléatoires. Lors de la planification d'un voyage sur Mars, il serait
pratique d'avoir des prix de billets à partir de plusieurs lignes spatiales en un seul endroit. Il existe des
sites Web qui regroupent les prix des billets pour les compagnies aériennes, mais rien jusqu'à présent pour les lignes
spatiales. Ce n’est cependant pas un problème pour vous. Vous pouvez utiliser Go pour apprendre à votre ordinateur
à résoudre des problèmes comme celui-ci. Commencez par construire un prototype qui génère 10 tickets aléatoires
et les affichés dans un format tabulaire avec un joli en-tête, comme suit:

Compagnies Spatiale	Jours		type voyage 	Prix
======================================================
Virgin Galactic 	23 			Round-trip		96 €
Virgin Galactic		39 			One-way			37 €
SpaceX				31			One-way			41 €
Space Adventures	22			Round-trip		100 €
Space Adventures	22			One-way			50 €
Virgin Galactic		30			Round-trip		84 €
Virgin Galactic		24			Round-trip		94 €
Space Adventures	27			One-way			44 €
Space Adventures	28			Round-trip		86 €
SpaceX				41			Round-trip		72 €

Le tableau doit avoir quatre colonnes:
- La compagnie spatiale fournissant le service
- La durée en jours du voyage vers Mars (aller simple)
- Si le prix couvre un aller-retour
- Le prix en millions d'Euros

Pour chaque billet, sélectionnez au hasard l'une des lignes spatiales suivantes: Space Adventures,
SpaceX ou Virgin Galactic. Utilisez le 13 octobre 2020 comme date de départ pour tous les billets.
Mars sera alors à 62 100 000 km de la Terre. Choisissez au hasard la vitesse à laquelle le navire
voyagera, de 16 à 30 km/s. Cela déterminera la durée du voyage vers Mars ainsi que le prix du billet.
Rendre les navires plus rapides le rendra du coup plus chers, dont le prix varie de 36 millions d'euros
à 50 millions d'euros. Doublez le prix des allers-retours.
*/

package main

import (
	"fmt"
	"math/rand"
)

const nombreSecondeParJour = 86400 // pour la convertion en jours

func main() {

	distanceTerreMars := 62100000
	compagnie := ""
	voyage := ""

	fmt.Println("Compagnies Spatiale	Jours		type voyage 		Prix")
	fmt.Println("====================================================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			compagnie = "Space Adventures"
		case 1:
			compagnie = "SpaceX"
		case 2:
			compagnie = "Virgin Galactic"
		}

		vitesse := rand.Intn(15) + 16
		duree := distanceTerreMars / vitesse / nombreSecondeParJour
		prix := 20 + vitesse

		if rand.Intn(2) == 1 {
			voyage = "Aller-Retour"
			prix *= 2
		} else {
			voyage = "Aller-simple"
		}

		fmt.Printf("%-21v %4v\t\t %-16v %8vM€\t\t\n", compagnie, duree, voyage, prix)
	}
}
