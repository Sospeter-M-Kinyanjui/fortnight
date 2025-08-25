package main

import "fmt"

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	result := ""
	single := ""
	count := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			single = zip(s[i], count)
			result += single
			single = ""
			count = 0
		} else {
			
		}
		count++
	}
	if count > 0 {
		result += single
	}
	return result
}

func zip(b byte, x int) string {
	result := ""
	result += itoa(x)
	for x > 0 {
		result += string(b)
		x--
	}
	return result
}

func itoa(x int) string {
	if x == 0 {
		return "0"
	}

	runes := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	digits := ""
	for x > 0 {
		digit := x % 10
		digits += string(runes[digit])
		x /= 10
	}
	return digits
}
