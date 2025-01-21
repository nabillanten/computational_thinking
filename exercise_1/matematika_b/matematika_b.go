package main

import "fmt"

func mathB(x float64) float64 {
	total := ((x * x * x) + 3*x) / ((x * x * x * x) - (3 * (x * x)) + 4)
	return total
}

func main() {
	var x float64

	fmt.Println("insert the x : ")
	fmt.Scanln(&x)

	result := mathB(x)

	fmt.Println("Output : ", result)
	fmt.Println("x = ", x)
}
