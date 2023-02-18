package main

import "fmt"

func main() {
	var n int
	fmt.Print("Program untuk membuat segitiga asterik")
	fmt.Print("Masukkan tinggi bintang	: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for z := 0; z <= i; z++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
