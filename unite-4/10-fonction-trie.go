package main

import (
	"fmt"
	"sort"
)

func main() {
	planete := []string{
		"Mercure", "Venus", "Terre", "Mars",
		"Neptune", "Uranus", "Saturne", "Jupiter"}

	sort.StringSlice(planete).Sort()
	fmt.Println(planete)
}
