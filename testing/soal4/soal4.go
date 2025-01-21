package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasiParkir, biaya int

	fmt.Scan(&jenisKendaraan, &durasiParkir)

	if jenisKendaraan == "Mobil" {
		if 0 < durasiParkir && durasiParkir <= 2 {
			biaya = durasiParkir * 3000
		} else if durasiParkir > 2 && durasiParkir <= 5 {
			biaya = durasiParkir*2000 + 3000
		} else if durasiParkir > 5 {
			biaya = durasiParkir*1000 + 3000
		}
	} else if jenisKendaraan == "Motor" {
		if 0 < durasiParkir && durasiParkir <= 2 {
			biaya = durasiParkir * 2000
		} else if durasiParkir > 2 && durasiParkir <= 5 {
			biaya = durasiParkir*1000 + 2000
		} else if durasiParkir > 5 {
			biaya = durasiParkir*500 + 2000
		}
	}

	fmt.Println(biaya)
}
