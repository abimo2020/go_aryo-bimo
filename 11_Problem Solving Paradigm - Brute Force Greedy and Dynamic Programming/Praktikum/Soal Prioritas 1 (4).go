package main

import (
	"fmt"
	"math"
)

func SimpleEquations(A, B, C int) {

	for y := 1; y <= 100; y++ {
		for z := 1; z <= 100; z++ {
			x := C - int(math.Pow(float64(z), 2)) - int(math.Pow(float64(y), 2))
			if x+y+z == A && x*y*z == B {
				fmt.Println(x, y, z)
				return
			}
		}

	}
	fmt.Println("No solution")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
