package main

import (
	"fmt"
)

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	f := food{}
	burger, chips, nuggets := 15, 10, 12
	if order == "burger" {
		f.preptime = burger
	} else if order == "chips" {
		f.preptime = chips
	} else if order == "nuggets" {
		f.preptime = nuggets
	}

	return f.preptime
}

func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
