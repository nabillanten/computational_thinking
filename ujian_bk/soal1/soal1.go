package main

import "fmt"

func main() {
	var firstDigit, secondDigit, thhirdDigit, fourthDigit, totalTwoNumber, totalThreeNumber, totalFirstAndLastNumber, factorNumber int

	fmt.Println("insert Vehicle Number")
	fmt.Scanln(&firstDigit, &secondDigit, &thhirdDigit, &fourthDigit)

	totalTwoNumber = firstDigit + secondDigit

	totalThreeNumber = secondDigit + thhirdDigit + fourthDigit

	totalFirstAndLastNumber = firstDigit + fourthDigit

	factorNumber = (secondDigit + thhirdDigit) % (totalFirstAndLastNumber)

	if totalTwoNumber == totalThreeNumber {
		println("Gerbang Ganjil")
	} else if factorNumber == 0 {
		println("Gerbang Genap")

	} else {
		println("Tidak Sama Sekali")
	}

}
