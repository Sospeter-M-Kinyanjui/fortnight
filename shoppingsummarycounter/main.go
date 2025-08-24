package main

import (
	"fmt"
)

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range shoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}

func shoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	item := ""

	for _, v := range str {
		if v == ' ' {
			result[item]++
			item = ""
			continue
		}
		item += string(v)
	}
	if len(item) > 0 {
		result[item]++
	}
	return result
}
