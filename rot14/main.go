package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func rot14(s string) string {
	result := ""

	for _, v := range s {
		// case for capital letters
		if v >= 65 && v <= 90 {
			temp := v + 14
			if temp > 90 {
				temp = temp - 90
				temp = (temp + 65) - 1
				result += string(temp)
			} else {
				result += string(temp)
			}
		// case for small letters
		} else if v >= 97 && v <= 122 {
			temp := v + 14
			if temp > 122 {
				temp = temp - 122
				temp = (temp + 97) - 1   // technically the code woun't work if you fail to -1
				result += string(temp)
			} else {
				result += string(temp)
			}
		// the rest of the characters to be appended as they are
		} else {
			result += string(v)
		}
	}
	return result
}
