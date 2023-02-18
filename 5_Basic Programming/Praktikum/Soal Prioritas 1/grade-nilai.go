package main

import "fmt"

func main() {
	var nilai int
	fmt.Println("Program untuk menentukan grade sebuah nilai")
	fmt.Print("Masukkan nilai	: ")
	fmt.Scan(&nilai)
	if nilai > 79 && nilai <= 100 {
		fmt.Println("Grade A")
	} else if nilai > 64 && nilai <= 79 {
		fmt.Println("Grade B")
	} else if nilai > 49 && nilai <= 64 {
		fmt.Println("Grade C")
	} else if nilai > 34 && nilai <= 49 {
		fmt.Println("Grade D")
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Println("Grade E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
