package main

import "fmt"

func main() {
	var number int

	fmt.Println("Insert Number")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Kelipatan 3")
		fmt.Println("Kelipatan 5")
	} else if number%3 == 0 {
		fmt.Println("Kelipatan 3")
	} else if number%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}
