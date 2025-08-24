package main

import (
	"fmt"
)

func main() {
	steps := collatzCountdown(12)
	fmt.Println(steps)
}

func collatzCountdown(start int) int {
	if start < 1 {
		return -1
	}
	count := 0
	for start != 1 {
		state := start % 2 == 0
		if state {
			start = start / 2
			count++
		} else {
			start = (start * 3) + 1
			count++
		}
	}
	return count
}
