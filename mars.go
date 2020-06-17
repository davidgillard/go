/*fichier : playground.go
auteur : David GILLARD */

/*
ce programme va permettre de connaître le poids et l'age de Nathan aurai si il vivait sur la planéte Mars

sachant qu'il pése 76kg et est agé de 45 ans

quelques explications:

- poids = force = newtons
- masse = kilogrammes
- poids = masse * accélération de la pesanteur, qui varie avec la planète


Sur Mars, on a la même masse, mais pas le même poids, car la gravité à la surface martienne est plus
faible. On pèse donc moins.
sur Terre je pèse 58 kg ; Pour calculer mon poids sur Mars :
Poids = masse × intensité de la pesanteur
Soit 58 × 3,7 = 214.6 N
soit 58 * 0.38 = 22.04 Kg
Sur Mars, mon poids est de 214.6 N et ma masse est de 58 kg
La planéte Mars prends 687 jours pour tournée autour du soleil
*/

package main // import du package main

import "fmt" // import de fmt du fournit les fonctions pour le formatage d'entrées / sorties

func main() { // le programme démarre avec la fonction du package main

	// Une première variante dans le programme consiste à écrire le code sur plusieurs ligne en répétant la fonction fmt.Println

	// fmt.Println("Mon poids à la surface de Mars serait de : ")
	// fmt.Println(75 * 0.38)
	// fmt.Println("kg et j'aurai ")
	// fmt.Println(45 * 365 / 687)
	// fmt.Println("ans.")

	// une autre façon de faire serait de toute faire sur une seule ligne
	// fmt.Println("Mon poids à la surface de Mars serait de", 75*0.38, "kg, et j'aurai", 45*365/687, "ans.")

	fmt.Printf("Mon poids à la surface de Mars serait de %v kg,", 75*0.38)
	fmt.Printf(" et j'aurai %v ans.\n", 45*365/687)
}
