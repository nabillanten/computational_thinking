package main

import (
	"fmt"
	"strconv"
)

func isPalindrom(num int) string {
	str := strconv.Itoa(num)
	strLength := len(str)

	for i := 0; i < strLength/2; i++ {
		if str[i] != str[strLength-1-i] {
			return "Bukan Palindrom"
		}
	}
	return "Palindrom"

}

func main() {

	var numbers int

	fmt.Println("Masukan angka")

	fmt.Scanln(&numbers)
	fmt.Println(isPalindrom(numbers))

}
