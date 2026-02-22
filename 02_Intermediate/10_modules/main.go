package main

// Import your custom module using the Module Name from `go.mod` + the folder path
import (
	"fmt"
	"myapp/mathutils"
)

func main() {
	// Call the public Add function from mathutils!
	sum := mathutils.Add(10, 20)
	fmt.Println("The sum is:", sum)

	// Will this line work? Why or why not?
	// result := mathutils.multiply(5, 5)
}
