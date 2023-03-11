package main

import (
	"fmt"
	"time"
)

func kelipatanN(n int) {
	i := 0
	for {
		if i%n == 0 {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
		i++
	}
}

func main() {
	go kelipatanN(7)
	time.Sleep(3 * time.Second)
}
