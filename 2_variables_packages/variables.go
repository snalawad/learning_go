package main

import "fmt"

const aConst int = 64

func main() {

	var aString string = "This is GO!!!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable type is %T\n", anotherString)

	newString := "This is also another string withhout declaring variables."
	fmt.Println(newString)
	fmt.Printf("The variable type is %T\n", newString)

	fmt.Printf("The const value %d\n", aConst)
}
