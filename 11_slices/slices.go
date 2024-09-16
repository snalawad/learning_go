package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)

	colors = append(colors, "purple")
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 43
	numbers[3] = 35
	numbers[4] = 23

	fmt.Println(numbers)

	numbers = append(numbers, 235)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
