package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		i := 0
		for {
			i++
			if i%3 == 0 {
				ch <- i
			}

		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
