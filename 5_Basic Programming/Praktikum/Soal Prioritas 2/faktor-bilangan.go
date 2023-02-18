package main

import "fmt"

func main() {
	var bilangan int
	fmt.Println("Program untuk menentukan faktor dari suatu bilangan")
	fmt.Print("Masukkan bilangan	: ")
	fmt.Scan(&bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
