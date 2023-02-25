package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayA = append(arrayA, arrayB...)
	result := []string{}
	condition := make(map[string]bool)

	for i := 0; i < len(arrayA); i++ {
		condition[arrayA[i]] = true
	}
	for key := range condition {
		result = append(result, key)
	}
	return result
}

func main() {
	arrayA := []string{"Aryo", "Bimo", "Aryo"}
	arrayB := []string{"Elsa", "Elsa", "Bimo", "Sabrina"}
	fmt.Println("Array A :", arrayA)
	fmt.Println("Array B :", arrayB)
	fmt.Println("Hasil Merge :", ArrayMerge(arrayA, arrayB))
}
