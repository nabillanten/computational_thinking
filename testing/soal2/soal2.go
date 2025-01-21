package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bilangan, jumlahSama int
	var stringBilangan string

	fmt.Scan(&bilangan)

	stringBilangan = strconv.Itoa(bilangan)

	for i := 0; i < len(stringBilangan); i++ {
		if stringBilangan[i] == stringBilangan[len(stringBilangan)-(i+1)] {
			jumlahSama++
		}
	}

	if jumlahSama == len(stringBilangan) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}

}
