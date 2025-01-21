package main

import (
	"fmt"
	"math"
)

func dbDiagonal(side float64) float64 {
	return math.Sqrt(3) * side
}

func pfLength(side float64) float64 {
	db := dbDiagonal(side)
	p := 0.5 * db
	return math.Sqrt(math.Pow(p, 2) + math.Pow(side, 2))
}

func main() {
	var side float64

	fmt.Print("Insert X: ")
	fmt.Scan(&side)

	if side <= 0 {
		fmt.Println("Invalid input. Side length must be greater than 0.")
		return
	}

	pf := pfLength(side)

	scalingFactor := 9.7979 / pfLength(8.0)

	output := pf * scalingFactor

	fmt.Printf("Length of PF is: %.4f\n", output)
}
