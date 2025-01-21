package main

import "fmt"

func main() {
	var result1, result2, result3, result4, result5 string
	var totalLost int

	fmt.Println("Insert 5 last match result (seperated with space)")
	fmt.Scanln(&result1, &result2, &result3, &result4, &result5)

	results := []string{result1, result2, result3, result4, result5}

	for _, result := range results {
		if result == "kalah" {
			totalLost++
		}
	}

	if totalLost >= 5 {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}

}
