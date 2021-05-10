package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	return kelvin(rand.Intn(151) + 150) // retourne un chiffre aléatoire de 151 à 301
}

func realSensor() kelvin {
	return 0
}

func main() {
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

}
