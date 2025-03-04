package main

import "fmt"

func main() {
	var k float64
	fmt.Scan(&k)

	fK:= (4*k+2) * (4*k+2) / ((4*k +1) * (4*k +3))
	fmt.Printf("nilai f(K) = %.10f\n", fK)
}