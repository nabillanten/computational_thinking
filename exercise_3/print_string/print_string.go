package main

import "fmt"

func main() {
	var n int
	var printedString string

	fmt.Println("Input number : ")
	fmt.Scanln(&n)

	fmt.Println("Input any string : ")
	fmt.Scanln(&printedString)

	fmt.Println("Result : ")

	for i := 0; i < n; i++ {
		fmt.Println(printedString)
	}
}
