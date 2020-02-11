package main

import (
	"strconv"
)

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Error converting into int")
	}
	return i
}

func validateCheckDigit(vin string) bool {
	total := 0

	checkDigit := stringToInt(string(vin[len(vin)-1]))

	for pos, char := range vin {
		if pos != len(vin)-1 {
			val := stringToInt(string(char))
			total += val * pos
		}
	}
	return total%11 == checkDigit
}

func vinNumbers(vin string) bool {
	if len(vin) == 10 {
		return true
	}
	return false
}

func main() {

}
