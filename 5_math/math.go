package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 54, 23
	intSum := i1 + i2 + i3

	fmt.Println("Integer Sum: ", intSum)

	f1, f2, f3 := 23.4, 45.3, 45.2
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float Sum rounded up:", floatSum)

	circleRadius := 15.5
	circumference := 2 * circleRadius * math.Pi
	fmt.Printf("Circumference of the circle: %.2f\n", circumference)
}
