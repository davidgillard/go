package main

import "fmt"

func main() {

	var distance int64 = 41.3e12
	const vitesseLumiere = 299792 // km/s
	const secondesEnJour = 86400

	fmt.Println("Alpha Centauri est à ", distance, "km de la planète Terre.")
	jours := distance / vitesseLumiere / secondesEnJour
	annees := jours / 365
	fmt.Println("C'est", jours, "jours de voyage à la vitesse de la lumière.")
	fmt.Println("soit", annees, "années à la vitesse de la lumière.")
}
