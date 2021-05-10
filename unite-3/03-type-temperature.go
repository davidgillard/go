package main

import "fmt"

type celsius float64
type farenheit float64
type kelvin float64

//kelvinToCelsius convertit une temperature en kelvin (°K) en celsius (°C)

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	var cel celsius = 127.0
	kel := celsiusToKelvin(cel)
	c := kelvinToCelsius(k)
	fmt.Print(k, "° K correspond à ", c, "° C\n")
	fmt.Print(cel, "° C correspond à ", kel, "° K\n")
}
