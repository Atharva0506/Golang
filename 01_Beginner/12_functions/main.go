package main

import "fmt"

func greetings(firstName string, lastName string) string {
	return "Welcome to Golang world, " + firstName + " " + lastName + "!"
}

// 2. Named Return Values
// Notice we defined `sum` and `totalVals` right in the signature.
// In Go, they automatically start as '0', so we don't need to initialize them!
func sumaAndLen(nums []int) (sum int, totalVals int) {
	for _, n := range nums {
		sum += n
		totalVals++
	}

	// 3. "Naked" or "Bare" Return
	// Since we named the return variables above, Go knows exactly what to return.
	return
}

// 4. Variadic Functions (...int)
func multiply(nums ...int) int {
	// Must start at 1, because 0 * anything = 0!
	mul := 1
	for _, n := range nums {
		mul *= n
	}
	return mul
}

func main() {
	// 1. Basic Function Call
	msg := greetings("Monkey D.", "Luffy")
	fmt.Println(msg)

	// 2. Calling a function with multiple return values
	mySlice := []int{10, 20, 30}
	totalSum, length := sumaAndLen(mySlice)
	fmt.Printf("Slice Sum: %v | Total Items: %v\n", totalSum, length)

	// 3. Calling a Variadic Function
	// We can pass as many ints as we want separated by commas!
	multipliedValue := multiply(2, 5, 10)
	fmt.Println("Multiplied Result (2 * 5 * 10):", multipliedValue)
}
