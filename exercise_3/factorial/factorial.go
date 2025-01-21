package main

import "fmt"

func main() {

	var numbers, result int

	result = 1

	fmt.Println("Input numbers to factorial : ")
	fmt.Scanln(&numbers)

	for i := 1; i <= numbers; i++ {
		result = result * i
	}

	fmt.Println("Factorial of", numbers, "is", result)
}
