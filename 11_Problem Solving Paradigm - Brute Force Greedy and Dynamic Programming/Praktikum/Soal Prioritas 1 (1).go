package main

import "fmt"

func biner(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%b \n", i)
	}
}

func main() {
	biner(5)
}
