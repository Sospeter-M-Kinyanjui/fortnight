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
	if len(slice1) == 0 && len(slice2) != 0 {
		return reverse(slice2)
	}
	if len(slice1) != 0 && len(slice2) == 0 {
		return reverse(slice1)
	}
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}

	var extra, longer, shorter, result []int
	if len(slice1) > len(slice2) {
		extra = slice1[len(slice2)-1:]
		shorter = slice2
		longer = slice1[:len(slice2)-1]
		extra, shorter, longer = reverse(extra), reverse(shorter), reverse(longer)
	} else if len(slice2) > len(slice1) {
		extra = slice2[len(slice1)-1:]
		shorter = slice1
		longer = slice2[:len(slice1)-1]
		extra, shorter, longer = reverse(extra), reverse(shorter), reverse(longer)
	} else {
		longer = slice1
		shorter = slice2
		longer, shorter = reverse(longer), reverse(shorter)
	}

	if len(slice1) != len(slice2) {
		result = append(result, extra...)
		for i := 0; i < len(longer); i++ {
			result = append(result, longer[i])
			result = append(result, shorter[i])
		}
	} else {
		for i := 0; i < len(longer); i++ {
			result = append(result, longer[i])
			result = append(result, shorter[i])
		}
	}

	return result
}

func reverse(slice []int) []int {
	result := []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}
