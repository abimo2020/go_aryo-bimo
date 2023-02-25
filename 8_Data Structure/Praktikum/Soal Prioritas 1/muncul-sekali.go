package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var result []int
	var condition = map[string]int{}

	for i := 0; i < len(angka); i++ {
		value := string(angka[i])
		var _, isExist = condition[value]
		if isExist {
			condition[value] = condition[value] + 1
		} else {
			condition[value] = 1
		}
		if condition[value] > 1 {
			delete(condition, value)
		}
	}
	for key := range condition {
		intKey, _ := strconv.Atoi(key)
		result = append(result, intKey)
	}
	return result
}

func main() {
	angka := "1234551"
	fmt.Println("Angka yang akan diproses adalah", angka)
	fmt.Println("Angka yang muncul sekali yaitu", munculSekali(angka))
}
