package main

import "fmt"

func main() {
	var alas, atap, tinggi int
	fmt.Println("Program untuk menghitung luas trapesium")
	fmt.Print("Masukkan panjang alas 	: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan panjang atap 	: ")
	fmt.Scan(&atap)
	fmt.Print("Masukkan panjang tinggi	: ")
	fmt.Scan(&tinggi)
	luas := (alas + atap) * tinggi / 2
	fmt.Println("Luas trapesium 		:", luas, "cm")
}
