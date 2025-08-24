package main

import "fmt"

func main() {
	fmt.Println(lastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(lastWord(" lorem,ipsum "))
	fmt.Println(lastWord(" "))
}

func lastWord(str string) string {
	words := trim(str)

	last := []byte{}

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == ' ' {
			break
		}
		last = append([]byte{words[i]}, last...)
	}
	return string(last)
}

func trim(str string) string {
	if len(str) == 1 {
		return str
	}
	start := 0
	end := len(str) - 1

	// trim leading
	for i := 0; i < len(str); i++ {
		if str[i] != 32 {
			break
		}
		start++
	}

	// adjust end
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != 32 {
			break
		}
		end--
	}

	return str[start : end+1]
}
