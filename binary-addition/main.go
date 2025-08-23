package main

import "fmt"

func main() {
	output := BinaryAddition(2, 2)
	fmt.Println(output)
}

func BinaryAddition(a int, b int) []int {
	if a < 0 || b < 0 {
		return nil
	}

	sum := a + b
	if sum == 0 {
		return []int{0}
	}

	result := []int{}
	for sum > 0 {
		bit := sum % 2
		result = append([]int{bit}, result...)
		sum /= 2
	}
	return result
}