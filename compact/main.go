package main

import (
	"fmt"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}

func compact(ptr *[]string) int {
	count := 0
	for i, v := range *ptr {
		if len(v) > 0 {
			count++
		} else if len(v) == 0 {
			if i < len(*prt) {
				*ptr = append((*ptr)[:i], (*ptr)[i+1:]...)
			}
		}
	}
	return count
}
