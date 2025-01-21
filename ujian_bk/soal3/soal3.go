package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("Masukan 3 angka x,y,z")
	fmt.Scanln(&x, &y, &z)

	numbers := []int{x, y, z}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	fmt.Println(numbers[0])

}
