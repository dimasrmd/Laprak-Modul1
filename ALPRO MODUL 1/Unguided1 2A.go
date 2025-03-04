package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)
	
	if tahun%400 == 0 {
		fmt.Println("Kabisat: true")
	} else if tahun%100 == 0 {
		fmt.Println("Kabisat: false")
	} else if tahun%4 == 0 {
		fmt.Println("Kabisat: true")
	} else {
		fmt.Println("Kabisat: false")
	}
}