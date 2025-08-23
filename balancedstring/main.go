package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	output := balancedString(args[0])
	printRunes(output)
}

func balancedString(arg string) int {
	balance := 0
	count := 0

	for _, v := range arg {
		if v == 'C' {
			balance++
		} else if v == 'D' {
			balance--
		}
		if balance == 0 {
			count++
		}
	}
	return count
}

func printRunes(x int) {
	if x == 0 {
		z01.PrintRune('0')
	}

	runes := []rune{'0','1','2','3','4','5','6','7','8','9'}
	for x > 0 {
		single := x % 10
		z01.PrintRune(runes[single])
		x /= 10
	}
	z01.PrintRune('\n')
}