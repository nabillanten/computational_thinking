package main

import (
	"fmt"
	"strconv"
)

func main() {
	var digit int
	var stringInput, firstDigit, lastDigit, doubleDigit string

	fmt.Println("Insert number : ")
	fmt.Scanln(&digit)

	stringInput = strconv.Itoa(digit)
	firstDigit = string(stringInput[0])
	lastDigit = string(stringInput[1])

	doubleDigit = firstDigit + firstDigit + lastDigit + lastDigit

	fmt.Println(doubleDigit)

}
