package main

import "fmt"

func main() {
	fmt.Println(firstWord("hello there"))
	fmt.Println(firstWord(""))
	fmt.Println(firstWord("hello   .........  bye"))
}

func firstWord(str string) string {
	words := trim(str)
	first := []byte{}

	for i := 0; i < len(words); i++ {
		if words[i] == ' ' {
			break
		}
		first = append(first, words[i])
	}
	return string(first)
}

func trim(str string) string {
	start := 0
	end := len(str) - 1

	// adjust start
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			start++
		} else {
			break
		}
	}

	// adjust end
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			end--
		} else {
			break
		}
	}

	return str[start : end+1]
}
