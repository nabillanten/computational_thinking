package main

import "fmt"

func main() {
	var amount, balance int

	fmt.Println("Input amount : ")
	fmt.Scanln(&amount)
	balance = amount
	for amount != 0 {
		fmt.Scanln(&amount)
		balance += amount
	}

	fmt.Println("your balance is :", balance)

}
