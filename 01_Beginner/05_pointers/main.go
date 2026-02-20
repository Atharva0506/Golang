package main

import "fmt"

func multiplyByTwo(num1 *int) {
	*num1 = *num1 * 2
}
func main() {

	num := "Batman"

	var ptr *string = &num // points to the address of num
	fmt.Println("Value of num: ", num)
	fmt.Println("Address of num: ", &num)
	fmt.Println("Value of ptr: ", ptr)
	fmt.Println("Address of ptr: ", &ptr)
	fmt.Println("Value of *ptr: ", *ptr)

	*ptr = "Superman" // changes the value of num
	fmt.Println("Value of num after change: ", num)
	fmt.Println("Address of num after change: ", &num)
	fmt.Println("Value of ptr after change: ", ptr)
	fmt.Println("Address of ptr after change: ", &ptr)
	fmt.Println("Value of *ptr after change: ", *ptr)

	// pointers with functions
	num2 := 10
	multiplyByTwo(&num2)
	fmt.Println("Value of num2 after function call: ", num2)
}
