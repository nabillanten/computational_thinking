package main

import "fmt"

func main() {

	var r float64
	var area int

	fmt.Println("Insert the radius (r) : ")
	fmt.Scanln(&r)

	area = int(4 * (22.0 / 7.0) * r * r)

	println("the output is :", area)
}
