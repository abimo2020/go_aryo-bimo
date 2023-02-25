package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	result := true
	if number < 2 {
		return false
	}
	// Bilangan prima akan memiliki sisa bagi 0 hingga pembaginya yaitu akar dari bilangannya sendiri kecuali bilangan itu sendiri
	akarN := int(math.Sqrt(float64(number)))
	for i := 2; i <= akarN; i++ {
		if number%i == 0 {
			result = false
			break
		}
	}
	return result
}

func main() {
	n := 13
	fmt.Println("Apakah", n, "merupakah bilangan prima?", primeNumber(n))
}
