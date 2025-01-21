package main

import "fmt"

func main() {
	var age int
	var haveAFamilyCard bool

	fmt.Println("Insert your age : ")
	fmt.Scanln(&age)
	fmt.Println("Do you have family card? :")
	fmt.Scanln(&haveAFamilyCard)

	if age >= 17 && haveAFamilyCard {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")

	}
}
