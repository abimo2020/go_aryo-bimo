package main

import "fmt"

func fibonanci(n int) int {
	fibo := []int{0, 1, 1}

	for i := 3; i <= n; i++ {
		fibo = append(fibo, fibo[i-1]+fibo[i-2])
	}

	return fibo[n]
}
func main() {
	fmt.Println(fibonanci(0))
	fmt.Println(fibonanci(1))
	fmt.Println(fibonanci(2))
	fmt.Println(fibonanci(3))
	fmt.Println(fibonanci(5))
	fmt.Println(fibonanci(6))
	fmt.Println(fibonanci(7))
	fmt.Println(fibonanci(9))
	fmt.Println(fibonanci(10))
}
