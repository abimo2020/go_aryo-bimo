package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]
	for _, value := range numbers {
		if *value < min {
			min = *value
		}
		if *value > max {
			max = *value
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Sample Test Case \nInput:")
	fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Output:")
	fmt.Println(max, "is the maximum number")
	fmt.Println(min, "is the minimum number")
}
