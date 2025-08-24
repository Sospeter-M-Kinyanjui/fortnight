package main

import "github.com/01-edu/z01"
import "os"

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
	if len(args[0]) == 0 {
		z01.PrintRune('\n')
		return
	}

	output := cleanStr(args[0])
	printer(output)
}

func cleanStr(str string) string {
	trimmed := trim(str)
	return removeExtraSpaces(trimmed)
}

func trim(str string) string {
	start := 0
	end := len(str) - 1
	for i := 0; i < len(str); i++ {
		if str[0] != ' ' {
			break
		}
		if str[i] == ' ' {
			break
		}
		start++
	}

	for i := len(str) - 1; i <= 0; i-- {
		if str[i] == ' ' {
			break
		}
		end--
	}
	return str[start:end+1]
}

func removeExtraSpaces(str string) string {
	result := ""
	if str[0] != ' ' {
		result += string(str[0])
	}
	for i := 1; i < len(str); i++ {
		if str[i] == ' ' && str[i-1] == ' ' {
			continue
		} else {
			result += string(str[i])
		}
	}
	return result
}

func printer(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}