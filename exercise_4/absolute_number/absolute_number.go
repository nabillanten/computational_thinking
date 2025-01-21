package main

import "fmt"

func main() {
	var number int
	fmt.Println("Insert Number")
	fmt.Scanln(&number)

	if number < 0 {
		number = number * (-1)
	}

	fmt.Println(number)
}
