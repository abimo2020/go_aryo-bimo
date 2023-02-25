package main

func pow(x, n int) int {
	result := x
	if n == 0 {
		result = 1
	} else if n == 1 {
		return result
	} else if n%2 == 0 {
		result *= x
	} else {
		result *= x * x
	}
	for i := 1; i < n/2; i++ {
		result *= x * x
	}
	return result
}

func main() {
	x := 2
	n := 1
	println("Hasil dari", x, "pangkat", n, ":", pow(x, n))
}
