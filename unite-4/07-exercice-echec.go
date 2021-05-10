package main

import "fmt"

func affichage(tableau [8][8]rune) {
	for _, ligne := range tableau {
		for _, colonne := range ligne {
			if colonne == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", colonne)
			}
		}
		fmt.Println()
	}
}

func main() {

	var tableau [8][8]rune
	tableau[0][0] = 'r'
	tableau[0][1] = 'n'
	tableau[0][2] = 'b'
	tableau[0][3] = 'q'
	tableau[0][4] = 'k'
	tableau[0][5] = 'b'
	tableau[0][6] = 'n'
	tableau[0][7] = 'r'

	for colonne := range tableau[1] {
		tableau[1][colonne] = 'p' //mise en place des pions pour les noirs
		tableau[6][colonne] = 'p' //mise en place des pions pour les blancs
	}

	tableau[7][0] = 'R'
	tableau[7][1] = 'N'
	tableau[7][2] = 'B'
	tableau[7][3] = 'Q'
	tableau[7][4] = 'K'
	tableau[7][5] = 'B'
	tableau[7][6] = 'N'
	tableau[7][7] = 'R'

	affichage(tableau)

}
