package main

import "fmt"

func main() {

	var vehicleType string
	var duration, totalFee int

	fmt.Println("Insert Vehicle Type and Duration")

	fmt.Scanln(&vehicleType, &duration)

	if vehicleType == "Mobil" {
		if duration <= 2 {
			totalFee = 3000
		} else if duration > 2 && duration <= 5 {
			totalFee = (2000 * duration) + 3000
		} else if duration > 5 {
			totalFee = (1000 * duration) + 3000
		}
	} else {
		if duration <= 2 {
			totalFee = 2000
		} else if duration > 2 && duration <= 5 {
			totalFee = (1000 * duration) + 2000
		} else if duration > 5 {
			totalFee = (500 * duration) + 2000
		}
	}

	fmt.Println(totalFee)

}
