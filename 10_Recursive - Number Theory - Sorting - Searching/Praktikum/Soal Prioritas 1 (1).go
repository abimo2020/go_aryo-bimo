package main

import "fmt"

func fibonanci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonanci(n-2) + fibonanci(n-1)
	}
}
func main() {
	fmt.Println(fibonanci(0))
	fmt.Println(fibonanci(2))
	fmt.Println(fibonanci(9))
	fmt.Println(fibonanci(10))
	fmt.Println(fibonanci(12))
}
