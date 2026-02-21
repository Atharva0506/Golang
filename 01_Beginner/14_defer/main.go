package main

import "fmt"

func openDb() {
	fmt.Println("Data Base connected...")
}
func dbClose() {
	fmt.Println("Database connection closed")
}
func main() {

	// 1. LIFO (Last In, First Out)
	// Because 3 was deferred last, it will execute first at the end of main()
	defer fmt.Println("Count 1")
	defer fmt.Println("Count 2")
	defer fmt.Println("Count 3")

	fmt.Println("Main execution started... (After 3 Defer Calls)")
	// 2. Resource Cleanup
	openDb()
	// We defer the close function IMMEDIATELY so we never forget it
	defer dbClose()

	// 3. Immediate Value Evaluation
	port := 80

	// This defer locks in the value '80' RIGHT NOW, even though it runs at the end
	defer fmt.Println("Port number inside defer: ", port)

	port = 8080
	fmt.Println("Port changed during execution: ", port)

}
