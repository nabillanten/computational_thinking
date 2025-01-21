package main

import "fmt"

func main() {
	var no1, no2, no3, no4 int

	fmt.Scan(&no1, &no2, &no3, &no4)

	if (no2+no3)%(no1+no4) == 0 {
		fmt.Println("Gerbang genap.")
	} else if no1+no2 == no2+no3+no4 {
		fmt.Println("Gerbang ganjil.")
	} else {
		fmt.Println("Tidak sama sekali.")
	}

}
