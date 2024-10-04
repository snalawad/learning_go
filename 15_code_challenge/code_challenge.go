package main

// Write your answer here, and then test your code.

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

func main() {
	slice := []string{"apple", "banana", "orange", "date"}
	result := convertToMap(slice)
	println(result)
}

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	// Your code goes here

	value := 100 / float64(len(items))

	for _, fruit := range items {
		result[fruit] = value
	}
	return result
}
