package main

import "fmt"

func main() {
	fmt.Print(fifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(fifthAndSkip("This is a short sentence"))
	fmt.Print(fifthAndSkip("1234"))
}

func fifthAndSkip(str string) string {
	if len(str) < 5 {
		return "Invalid Input\n"
	} else if len(str) == 0 {
		return "\n"
	}
	result := ""
	count := 0
	foo := removeSpaces(str)
	for _, v := range foo {
		if count == 5 {
			result += " "
			count = 0
			continue
		}
		result += string(v)
		count++
	}
	result += "\n"
	return result
}

func removeSpaces(str string) string {
	result := ""
	for _, v := range str {
		if v != ' ' {
			result += string(v)
		}
	}
	return result
}
