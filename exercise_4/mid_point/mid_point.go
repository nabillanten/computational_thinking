package main

import (
	"fmt"
)

func main() {
	var midPoint, startNumber, endNumber int
	fmt.Println("Insert three numbers (seperated with space)")
	fmt.Scanln(&midPoint, &endNumber, &startNumber)

	if ((startNumber + endNumber) / 2) == midPoint {
		fmt.Println(true)
		fmt.Println("Penjelasan", midPoint, "adalah nilai tengah dari ", startNumber, "Sampai", endNumber)
	} else {
		fmt.Println(false)
	}

}
