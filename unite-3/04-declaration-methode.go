/*
Write a program with celsius, fahrenheit, and kelvin types and methods to convert from
any temperature type to any other temperature type.
*/

package main

import (
	"fmt"
)

// definition du type celsius

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// définition du type fahrenheit
type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// imcompréhension
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// imcompréhension
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Print(k, "° K correspond à ", c, "° C\n")
}
