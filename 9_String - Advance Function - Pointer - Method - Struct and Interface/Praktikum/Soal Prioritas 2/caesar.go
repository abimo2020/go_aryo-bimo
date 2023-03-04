package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	result := []rune{}
	for _, value := range input {
		value = 'a' + (value-'a'+rune(offset))%26
		result = append(result, value)
	}
	return string(result)
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
