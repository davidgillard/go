package main

import "fmt"

func main() {

	var planete [8]string
	planete[0] = "Mercure"
	planete[1] = "Venus"
	planete[2] = "Terre"

	terre := planete[2]
	fmt.Println(terre)
	fmt.Println(len(planete))
	fmt.Println(planete[3] == "") // ici on verifie que le 4iéme élément du tableau est vide
	//	planete[8] = "Pluton"
	//	pluton := planete[8]
	// Une façon qu'offre Go pour déclarer un tableau
	dwarfs := [5]string{"Меркурий", "Плутон", "Венера", "маршировать", "суша"}
	fmt.Println(dwarfs)

}
