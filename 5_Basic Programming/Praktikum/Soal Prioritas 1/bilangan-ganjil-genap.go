package main

import "fmt"

func main() {
	var bilangan int
	fmt.Println("Program untuk menentukan bilangan ganjil/genap")
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bilangan)
	if bilangan%2 == 0 {
		fmt.Println("Bilangan tersebut adalah bilangan genap")
	} else {
		fmt.Println("Bilangan tersebut adalah bilangan ganjil")
	}
}
