package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("0")
		return
	}

	sum := addPrimeSum(args[0])
	fmt.Println(sum)
}

func addPrimeSum(str string) int {
	endpoint := atoi(str)

	if endpoint <= 1 {
		return endpoint
	}
	sums := 0
	for i := 2; i <= endpoint; i++ {
		if isPrime(i) {
			sums += i
		}
	}
	return sums
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func atoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	isNeg := false
	if str[0] == '-' {
		isNeg = true
	}
	digits := 0
	for _, v := range str {
		if v >= 48 && v <= 57 {
			digit := int(v - '0')
			digits += (digits * 10) + digit
		} else {
			return 0
		}
	}
	if isNeg {
		return digits * -1
	}
	return digits
}
