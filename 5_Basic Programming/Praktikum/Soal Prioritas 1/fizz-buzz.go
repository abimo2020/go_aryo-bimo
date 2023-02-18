package main

import "fmt"

func main() {
	fmt.Println("Program mencetak angka 1-100 Fizz Buzz")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("fizz ")
		} else if i%5 == 0 {
			fmt.Print("buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
