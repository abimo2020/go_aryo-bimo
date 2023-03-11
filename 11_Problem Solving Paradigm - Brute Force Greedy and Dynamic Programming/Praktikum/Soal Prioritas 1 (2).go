package main

import "fmt"

func pascal(numRows int) [][]int {
	result := [][]int{}
	bilangan := []int{1}

	for i := 0; i < numRows; i++ {
		result = append(result, append([]int(nil), bilangan...))
		for j := i; j > 0; j-- {
			bilangan[j] += bilangan[j-1]
		}
		bilangan = append(bilangan, 1)
	}
	return result
}

func main() {
	fmt.Println(pascal(5))
}
