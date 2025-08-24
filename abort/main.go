package main

import "github.com/01-edu/z01"

func main() {
	result := Abort(2, 3, 8, 5, 7)
	printer(result)
}

func printer(x int) {
	if x == 0 {
		z01.PrintRune('0')
	}

	runes := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for x > 0 {
		digit := x % 10
		z01.PrintRune(runes[digit])
		x /= 10
	}
	z01.PrintRune('\n')
}

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	sorted := quick(slice)
	firsthalf := sorted[:len(sorted)/2]
	secondhalf := sorted[len(sorted)/2:]
	if len(firsthalf) > len(secondhalf) {
		return firsthalf[len(firsthalf)-1]
	} else if len(secondhalf) > len(firsthalf) {
		return secondhalf[0]
	} else {
		temp := firsthalf[len(firsthalf)-1] + secondhalf[0]
		return temp / 2
	}
}

func quick(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[len(array)/2]
	var left, middle, right []int

	for _, v := range array {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}

	sortedLeft := quick(left)
	sortedRight := quick(right)

	final := sortedLeft
	final = append(final, middle...)
	final = append(final, sortedRight...)

	return final
}
