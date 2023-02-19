package main

import (
	"fmt"
)

func main() {
	var kata string
	hasil := "Palindrome"

	fmt.Println("Apakah Palindrome ?")
	fmt.Print("masukan kata: ")
	fmt.Scanln(&kata)

	length := len(kata)

	for i := 0; i < length/2; i++ {
		if kata[i] != kata[length-1-i] {
			hasil = "Bukan Palindrome"
			break
		}
	}

	fmt.Println("captured:", kata)
	fmt.Println(hasil)
}
