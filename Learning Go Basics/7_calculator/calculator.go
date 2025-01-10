// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Print(err)
		panic("Value 1 must be a number")
	}

	// Convert the second string to a float64
	bFloat, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Print(err)
		panic("Value 2 must be a number")
	}
	// Calculate and return the result

	sumValue := aFloat + bFloat

	return sumValue
}

func main() {

	value1 := "100.0"
	value2 := "-50.5"
	result := calculate(value1, value2)
	fmt.Println("Result: ", result)
}
