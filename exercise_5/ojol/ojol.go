package main

import (
	"fmt"
)

func main() {
	var hour, minute int
	var distance, ratePerKm float64

	fmt.Println("Insert hour, minute, and distance (seperated with space)")

	fmt.Scanln(&hour, &minute, &distance)

	if hour < 5 || hour >= 22 || (hour == 22 && minute > 0) {
		fmt.Println("Maaf, kami tidak bisa melayani pesanan anda")
	}

	if (hour >= 5 && hour < 9) || (hour >= 16 && hour < 19) {
		if distance <= 10 {
			ratePerKm = 5000
		} else if distance <= 20 {
			ratePerKm = 4500
		}
	} else {
		if distance <= 20 {
			ratePerKm = 4000
		}
	}

	totalPrice := ratePerKm * distance

	fmt.Println(totalPrice)
}
