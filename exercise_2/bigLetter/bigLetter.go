package main

import (
	"fmt"
	"regexp"
)

func main() {
	var character string
	fmt.Println("Insert the character : ")
	fmt.Scanln(&character)

	largeRegexp := `^[A-Z]$`

	isBigLetter := regexp.MustCompile(largeRegexp)

	fmt.Println(isBigLetter.MatchString(character))

}
