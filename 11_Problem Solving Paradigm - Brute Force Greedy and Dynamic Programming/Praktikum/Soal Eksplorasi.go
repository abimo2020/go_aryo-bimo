package main

import "fmt"

func romawi(n int) string {
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	cost := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for i := 0; i < len(romawi); i++ {
		for n >= cost[i] {
			n -= cost[i]
			result += romawi[i]
		}
	}
	return result
}

func main() {
	fmt.Println(romawi(4))
	fmt.Println(romawi(9))
	fmt.Println(romawi(23))
	fmt.Println(romawi(2021))
	fmt.Println(romawi(1646))
}
