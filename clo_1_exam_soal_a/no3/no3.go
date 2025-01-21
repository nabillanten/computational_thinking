package main

import "fmt"

func main() {

	var number int

	fmt.Println("Insert Number :")
	fmt.Scanln(&number)

	if number < 6 {
		fmt.Println(false)
		return
	}

	for number >= 6 {
		number = number - 6
	}

	if number == 0 {
		fmt.Println(true)
		return
	} else {
		fmt.Println(false)
	}
}
