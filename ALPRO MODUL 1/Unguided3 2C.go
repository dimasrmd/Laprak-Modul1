package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa_gram := berat % 1000

	biaya := kg * 10000
	biaya_tambahan := 0

	switch {
	case sisa_gram == 0:
		biaya_tambahan = 0
	case sisa_gram >= 500:
		biaya_tambahan = sisa_gram * 5
	case sisa_gram < 500:
		biaya_tambahan = sisa_gram * 15
	}
	if kg >= 10 {
		biaya_tambahan = 0
	}

	totalBiaya := biaya + biaya_tambahan

	fmt.Println("Detail berat:", kg, "kg +", sisa_gram, "gr")
	fmt.Println("Detail biaya: Rp.", biaya, "+ Rp.", biaya_tambahan)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
