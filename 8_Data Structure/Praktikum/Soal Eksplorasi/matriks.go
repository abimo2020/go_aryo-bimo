package main

import "math"

func main() {
	matriks := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	diagonal1 := 0
	diagonal2 := 0
	length := len(matriks)

	for i := 0; i < length; i++ {
		diagonal1 += matriks[i][i]
		diagonal2 += matriks[i][length-i-1]
	}
	result := int(math.Abs(float64(diagonal1 - diagonal2)))

	println("Total dari diagonal 1 :", diagonal1)
	println("Total dari diagonal 2 :", diagonal2)
	println("Perbedaan mutlak dari kedua diagonal :", result)
}
