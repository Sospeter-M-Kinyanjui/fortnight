package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Not enough arguments")
		return
	}

	number := atoi(args[0])
	fmt.Println(number)
}

func atoi(str string) int {
	number := 0
	for _, v := range str {
		if v >= '0' && v <= '9' {
			digit := int(v - '0')
			number = (number * 10) + digit
		} else {
			return -1
		}
	}
	return number
}

// "857"
