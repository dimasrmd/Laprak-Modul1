package main

import "fmt"

func main() {
	var greetings = "Selamat datang di dunia DAP"
	var a, b int
	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}

//package main

//import "fmt"

//func main() {
	//for i := 1; i <= 5; i++ {
		//fmt.Println("Iterasi ke-", i)
		
	//}
//}