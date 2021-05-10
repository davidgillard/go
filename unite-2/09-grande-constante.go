package main

import (
	"fmt"
)

func main() {

	const distance = 236000000000000000
	const joursEnSecondes = 86400
	const vitesseLumiere = 299792
	const joursParAn = 365

	const annees = distance / vitesseLumiere / joursEnSecondes / joursParAn
	fmt.Println("la galaxie est à ", annees, "années lumière de la Terre")
}
