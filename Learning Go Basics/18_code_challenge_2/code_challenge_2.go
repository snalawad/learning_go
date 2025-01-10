// Write your answer here, and then test your code.

package main

import (
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {

	var total_value float64

	for _, cart_item := range cart {
		fmt.Println(cart_item)
		total_value += float64(cart_item.price) * float64(cart_item.quantity)
	}

	fmt.Println(total_value)
	return 0
}

func main() {

	// You can edit this code to try different testing cases.
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Println(result)
}
