package main

import (
	"fmt"
)

func isPrime(number, i int) bool {
	if number == i {
		return true
	} else if number%i == 0 {
		return false
	} else {
		return isPrime(number, i+1)
	}
}
func primeX(number int) int {
	result := []int{}
	if number < 1 {
		return 0
	} else {
		for i := 2; i <= number*100; i++ {
			if isPrime(i, 2) == true {
				result = append(result, i)
				if len(result) == number {
					break
				}
			}
		}
	}

	return result[number-1]
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
