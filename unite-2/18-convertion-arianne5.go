package main

import (
	"fmt"
	"math"
)

func main() {

	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println("La valeur de bh après convertion en int16", h)
	fmt.Println("la valeur min de la convertion en int16 est:", math.MinInt16)
	fmt.Println("la valeur max de la convertion en int16 est:", math.MaxInt16)
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("Veuillez préparer vos bagages")
	}
}
