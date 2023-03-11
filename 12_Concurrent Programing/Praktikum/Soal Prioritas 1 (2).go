package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			i++
			if i%3 == 0 {
				ch <- i
			}

		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			number := <-ch
			fmt.Println(number)
		}

	}()
	<-time.After(time.Second * 1)
}
