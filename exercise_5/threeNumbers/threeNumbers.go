package main

import "fmt"

func main() {
	var firstNumber, secondNumber, thirdNumber int
	fmt.Println("Insert three numbers (seperated with space)")
	fmt.Scanln(&firstNumber, &secondNumber, &thirdNumber)

	numbers := []int{firstNumber, secondNumber, thirdNumber}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	fmt.Println(numbers[0], numbers[1], numbers[2])

}
