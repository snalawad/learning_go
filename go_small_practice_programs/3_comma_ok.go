package main

import (
	"fmt"
)

func main() {
	prices := map[string]int{
		"Banana": 0,
	}

	price, ok := prices["Banana"]
	if ok {
		fmt.Printf("The price of Banana is $%d\n", price)
	} else {
		fmt.Printf("We Don't have Bananas")
	}

	price, ok = prices["Apple"]
	if ok {
		fmt.Printf("The price of Apple is $%d\n", price)
	} else {
		fmt.Printf("We Don't have Apples")
	}
}
