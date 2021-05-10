package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFarenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFarenheit(k float64) float64 {
	return celsiusToFarenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "° K est ", celsius, "° C\n")
	fmt.Printf("233° K correspond à %2.f° C\n", kelvinToCelsius(233))
	fmt.Printf("0° K correspond à %2.f° F", kelvinToFarenheit(0))
}
