package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}

func Unmatch(a []int) int {
	matches := make(map[int]int)

	for _, v := range a {
		matches[v]++
	}

	for k, v := range matches {
		if v == 1 {
			return k
		}
	}
	return -1
}
