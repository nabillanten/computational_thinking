package main

import "fmt"

func main() {
	var n, m, x, y, cups int
	fmt.Println("Insert n, m, x, y (seperated with space)")
	fmt.Scanln(&n, &m, &x, &y)

	for n >= x && m >= y {
		n -= x
		m -= y
		cups++
	}

	fmt.Println(cups)

}
