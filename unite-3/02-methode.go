package main

import (
	"fmt"
)

func main() {

	type celsius float64
	const degree = 20
	var temperature celsius = degree
	temperature += 10
	fmt.Println(temperature)
}
