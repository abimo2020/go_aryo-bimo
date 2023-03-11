package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, value := range cards {
		if value[0] == deck[0] || value[0] == deck[1] || value[1] == deck[0] || value[1] == deck[1] {
			return value
		}
	}
	return "Tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 11}))
}
