package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// rand.NewSource(time.Now().Unix())
	// print(args ...Type)
	var result string

	switch dow := rand.Int() + 1; dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)
}
