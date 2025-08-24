package main

import "fmt"

func main() {
	arg1 := 4
	fmt.Println(fibonacci(arg1))
}

func fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index <= 1 {
			return index
	}
	return fibonacci(index-1) + fibonacci(index-2)
}
