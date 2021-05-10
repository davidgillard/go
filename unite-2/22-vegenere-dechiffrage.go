package main

func main() {

	texteChiffre := "CSOITEUIWUIZNSROCNKFD"
	motClef := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(texteChiffre); i++ {
		c := texteChiffre[i] - 'A'
		k := motClef[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'

		message += string(c)

		keyIndex++
	}
}
