package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Insert three sides of triangle (seperated with space)")
	fmt.Scanln(&a, &b, &c)

	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == c || b == c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}
