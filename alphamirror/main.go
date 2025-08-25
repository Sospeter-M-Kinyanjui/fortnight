package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	result := alphaMirror((args[0]))
	fmt.Println(string(result))
}

func alphaMirror(str string) []rune {
	slice := []rune(str)

	for i, v := range slice {
		if v >= 'a' && v <= 'z' {
			oposite := 'a' + ('z' - v)
			slice[i] = oposite
		} else if v >= 'A' && v <= 'Z' {
			oposite := 'A' + ('Z' - v)
			slice[i] = oposite
		}
	}
	return slice
}
