package main

import "fmt"

func main() {
	var unit uint8 = 255
	fmt.Println(unit)

	var myInt int = 10
	fmt.Println(myInt)

	var myFloat float64 = 10.5
	fmt.Println(myFloat)

	var myStr string = "hey luffy!!"
	fmt.Println(myStr)

	var myBool bool = true
	fmt.Println("bool: ", myBool)

	// Type inference
	var num = 10 // int automatically detect the type
	fmt.Printf("Type: %T Value: %v\n", num, num)

	// Short variable declaration
	name := "luffy" //can only be used inside a function
	fmt.Println(name)

	// Constants
	const pi = 3.14
	fmt.Println(pi)

	// multiple variables
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// multiple variables with type inference
	d, e, f := 4, 5, 6
	fmt.Println(d, e, f)
}
