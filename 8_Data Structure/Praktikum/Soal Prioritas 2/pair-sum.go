package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var result = []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = append(result, arr[i])
				result = append(result, arr[j])
			}
		}
		if len(result) == 2 {
			break
		}
	}
	return result
}

func main() {
	array := []int{1, 3, 5, 9, 13, 14, 15}

	target := 16
	fmt.Println("Array :", array)
	fmt.Println("Target :", target)
	fmt.Println("Pair sum :", PairSum(array, target))

}
