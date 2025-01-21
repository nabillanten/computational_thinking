package main

import "fmt"

func main() {
	var x, y, z, buffer int

	fmt.Scan(&x, &y, &z)

	if x > y {
		buffer = x
		x = y
		y = buffer
	}

	if y > z {
		buffer = y
		y = z
		z = buffer
	}

	if x > y {
		buffer = x
		x = y
		y = buffer
	}

	fmt.Printf("%d, Penjelasan: Bilangan paling besar adalah %d", z, z)
}
