package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("The current time is: ", n)

	t := time.Date(2009, time.August, 11, 8, 43, 12, 23, time.UTC)
	fmt.Println("The date is: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:01:3 2009")
	fmt.Printf("The type of parsedTime is: %T\n", parsedTime)

}
