package main

import "fmt"

func mathA(x float64, y float64) float64 {
	total := (1.0 / (3.0*x*x + 10)) + ((10 * y) + 7)
	return total
}

func main() {
	var x, y float64

	fmt.Println("insert the x : ")
	fmt.Scanln(&x)

	fmt.Println("insert the y : ")
	fmt.Scanln(&y)

	result := mathA(x, y)

	fmt.Println("Output : ", result)
	fmt.Println("x = ", x, "dan y = ", y)
}
