package main

import (
	"fmt"
)

func main() {
	fmt.Println(camelToSnakeCase("HelloWorld"))
	fmt.Println(camelToSnakeCase("helloWorld"))
	fmt.Println(camelToSnakeCase("camelCase"))
	fmt.Println(camelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(camelToSnakeCase("camelToSnakeCase"))
	fmt.Println(camelToSnakeCase("hey2"))
}

func camelToSnakeCase(str string) string {
	if str[len(str)-1] >= 65 && str[len(str)-1] <= 90 {
		return str
	}
	result := ""

	result += string(str[0])
	for i := 1; i < len(str); i++ {
		if (str[i] >= 65 && str[i] <= 90) && (str[i-1] >= 90 && str[i-1] <= 122) {
			result += "_"
			// foo := string(str[i] + 32)
			result += string(str[i])
		} else {
			result += string(str[i])
		}
	}
	return result
}