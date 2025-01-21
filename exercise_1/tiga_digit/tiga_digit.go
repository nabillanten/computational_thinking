package main

import "fmt"

type Result struct {
	d1        int
	d2        int
	d3        int
	intConcat string
}

func tiga_digit(numbers int) Result {
	var d1, d2, d3 int

	d1 = numbers / 100
	d1 = d1 % 10

	d2 = numbers / 10
	d2 = d2 % 10

	d3 = numbers % 10

	result := Result{
		d1:        d1,
		d2:        d2,
		d3:        d3,
		intConcat: fmt.Sprintf("%d %d %d", d1, d2, d3),
	}

	return result
}

func main() {
	var x int

	fmt.Println("insert the x : ")
	fmt.Scanln(&x)

	result := tiga_digit(x)

	fmt.Println("Output : ", result.intConcat)
	fmt.Println("x =", x, "maka d1 =", result.d1, ",d2 =", result.d2, ",d3 =", result.d3)
}
