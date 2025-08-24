package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	forEach(printNbr, a)
	fmt.Println()
}

func forEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

func printNbr(x int) {
	fmt.Print(x)
}
