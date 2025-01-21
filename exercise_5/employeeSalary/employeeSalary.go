package main

import "fmt"

func main() {
	var position string
	var yearOfService, childAllowance, salary, totalChild int

	fmt.Println("Insert position, year of service, and total child (seperated with space)")
	fmt.Scanln(&position, &yearOfService, &totalChild)

	if totalChild >= 3 {
		totalChild = 3
	}

	if position == "Staf" {
		if yearOfService < 5 {
			salary = 4000
		} else if yearOfService >= 5 && yearOfService <= 10 {
			salary = 4000
			childAllowance = 100 * totalChild
		} else {
			salary = 5000
			childAllowance = 100 * totalChild
		}
	} else if position == "Manager" {
		if yearOfService <= 10 {
			salary = 8500
		} else {
			salary = 10000
		}
		childAllowance = 300 * totalChild
	} else {
		salary = 20000
		childAllowance = 500 * totalChild
	}
	fmt.Println(salary, "+", childAllowance, "=", salary+childAllowance)
}
