package main

import "fmt"

func main() {

	var numbers string
	var total_number int

	fmt.Println("Insert Number :")
	fmt.Scanln(&numbers)

	for _, s := range numbers {
		if string(s) == "1" {
			total_number += 1
		}
	}
	fmt.Println(total_number)
}
