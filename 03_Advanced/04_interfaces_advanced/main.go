package main

import "fmt"

func ProcessData(data any) {
	// The Type Switch: This safely unpacks the `any` interface.
	// `v` instantly becomes the physical underlying type (int, string, etc) inside the case block!
	switch v := data.(type) {
	case int:
		// Because it matched `int`, the compiler now lets us do math on `v`!
		fmt.Printf("%d\n", v*2)
	case string:
		fmt.Println(v)
	case bool:
		// Used Println instead of Print so the next line doesn't start on the same line in the terminal!
		fmt.Println("Boolean Value:", v)
	default:
		// Default catches anything we didn't explicitly check for (like your float64!)
		fmt.Println("Unknown data type")
	}
}
func main() {
	ProcessData(12)
	ProcessData("Yo")
	ProcessData(true)
	ProcessData(12.11)
}
