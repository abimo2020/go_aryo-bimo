package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var result = map[string]int{}

	for i := 0; i < len(slice); i++ {
		var _, isExist = result[slice[i]]
		if isExist {
			result[slice[i]] = result[slice[i]] + 1
		} else {
			result[slice[i]] = 1
		}
	}
	return result
}

func main() {
	slice := []string{"a", "b", "c", "d", "a", "b", "a", "c"}
	fmt.Println("Slice sebelum mapping :", slice)
	fmt.Println(Mapping(slice))
}
