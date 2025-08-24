package main

import (
	"fmt"
)

func main() {
	chunk([]int{}, 10)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}

	chunks := [][]int{}
	single := []int{}

	counter := 0

	for _, v := range slice {
		if counter == size {
			chunks = append(chunks, single)
			single = nil
			counter = 0
		}
		single = append(single, v)
		counter++
	}
	if counter > 0 {
		chunks = append(chunks, single)
	}
	fmt.Println(chunks)
}
