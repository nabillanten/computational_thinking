package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var char, lowerCaseChar string
	consonant := []string{"a", "i", "u", "e", "o"}
	fmt.Println("Insert Character")
	fmt.Scanln(&char)

	lowerCaseChar = strings.ToLower(char)

	if !slices.Contains(consonant, lowerCaseChar) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan konsonan")
	}
}
