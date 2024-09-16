package main

import (
	"fmt"
)

func main() {
	// Create an array noramlly by declaring it first and the
	// add the values into it.

	var colors [3]string

	colors[0] = "Red"
	colors[1] = "Yellow"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println(colors[2])

	// Create an array in a single line with type of data, length of array
	// and the data

	var numbers = [5]int{5, 3, 1, 8, 9}
	fmt.Println(numbers)

	fmt.Println("Count of colors: ", len(colors))
	fmt.Println("Count of numbers: ", len(numbers))

}
