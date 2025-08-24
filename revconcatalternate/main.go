package main

import (
	"fmt"
)

func main() {
	fmt.Println(revConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(revConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(revConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(revConcatAlternate([]int{1, 2, 3}, []int{}))
}

func revConcatAlternate(slice1 []int, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)

	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	result := make([]int, 0, len1+len2)

	for i := maxLen - 1; i >= 0; i-- {
		if i < len1 {
			result = append(result, slice1[i])
		}
		if i < len2 {
			result = append(result, slice2[i])
		}
	}

	return result
}
