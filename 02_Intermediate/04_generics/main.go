package main

import "fmt"

type Number interface {
	int | float64
}

// 1. Generic Sum function using our Number constraint
func Sum[T Number](nums []T) T {
	// Job Tip: You can't just write `total := 0` because 0 is an integer, and `T` might become a `float64`!
	// Instead, we declare it as type `T`. Go will automatically set it to the zero-value (0 or 0.0) for us!
	var total T

	for _, num := range nums {
		total += num
	}
	return total
}

// 2. Generic Print function that accepts absolutely 'any' type
func PrintArray[T any](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
func main() {
	num := []int{1, 22, 3}
	words := []string{"Luffy", "Is", "One", "Picec", "Character"}

	// Calling the ANY generic function
	fmt.Print("Numbers Array: ")
	PrintArray(num)

	fmt.Print("Words Array: ")
	PrintArray(words)

	// Calling the NUMBER generic function (Testing constraint with ints and floats)
	floatNums := []float64{10.5, 20.5}

	fmt.Println("Sum of Ints:", Sum(num))         // Uses int math
	fmt.Println("Sum of Floats:", Sum(floatNums)) // Uses float64 math
}
