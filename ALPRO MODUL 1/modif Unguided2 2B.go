package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	result := 1.0
	for i := 0; i <= k; i++ {
		pembilang := (4*i + 2) * (4*i + 2)
		penyebut := (4*i + 1) * (4*i + 3)
		result *= float64(pembilang) / float64(penyebut)
	}
	fmt.Printf("nilai √2 = %.10f\n", result)
}
