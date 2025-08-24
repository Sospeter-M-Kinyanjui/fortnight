package main

import (
	"fmt"
)

func main() {
	fmt.Println(concatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(concatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(concatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(concatAlternate([]int{1, 2, 3}, []int{}))
}

func concatAlternate(slice1 []int, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) != 0{
		return slice2
	}
	if len(slice1) != 0 && len(slice2) == 0 {
		return slice1
	}
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}
	var extra, longer, shorter, result []int
	if len(slice1) > len(slice2) {
		extra = slice1[len(slice2)-1:]
		shorter = slice2
		longer = slice1[:len(slice2)-1]
	} else if len(slice2) > len(slice1) {
		extra = slice2[len(slice1)-1:]
		shorter = slice1
		longer = slice2[:len(slice1)-1]
	} else {
		longer = slice1
		shorter = slice2
	}

	for i := 0; i < len(longer); i++ {
		result = append(result, longer[i])
		result = append(result, shorter[i])
	}
	result = append(result, extra...)

	return result
}

