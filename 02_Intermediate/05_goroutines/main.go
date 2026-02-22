package main

import (
	"fmt"
	"time"
)

// 1. Basic Named Goroutine
func fetchData(source string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Printf("Data Fetched from: %s\n", source)
}

func main() {
	// These run concurrently!
	go fetchData("Database", 2*time.Second) // Takes 2 seconds
	go fetchData("API", 1*time.Second)      // Takes 1 second (Will finish FIRST!)
	// 2. Anonymous Goroutine
	// This spins up a background task instantly without needing a named function!
	go func(port int) {
		time.Sleep(1 * time.Second)
		fmt.Printf("Server Started at PORT: %d\n", port)
	}(8000)

	// Since we launched all three Goroutines at the exact same time,
	// we only need to sleep long enough to let the SLOWEST one finish (Database at 2 seconds).
	// Let's sleep for exactly 3 seconds to be safe.
	time.Sleep(3 * time.Second)

	fmt.Println("Main function finished, shutting down all remaining Goroutines!")
}
