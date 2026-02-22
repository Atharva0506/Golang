package main

import (
	"fmt"
	"time"
)

// 1. Basic Channel Usage
func calculateSquare(number int, resultChan chan int) {
	num := number * number
	resultChan <- num // Push the data INTO the channel pipe
}

// 2. Real World Example: Fetching from multiple APIs concurrently
func fetchPrice(apiName string, delay time.Duration, priceChan chan string) {
	// Simulate an HTTP request that takes "delay" seconds
	time.Sleep(delay)

	// When the HTTP request finishes, push the string into the channel!
	priceChan <- fmt.Sprintf("API Result from %s", apiName)
}
func main() {
	// ==== Part 1: Basic Channels ====
	numChannel := make(chan int)
	go calculateSquare(5, numChannel)

	// This `<-` operator physically FREEZES main() until the data arrives!
	result := <-numChannel
	fmt.Println("Squared Result:", result)

	// ==== Part 2: Real World Concurrency (Like Promise.all in JS) ====
	fmt.Println("\n--- Starting API Fetch ---")

	priceChan := make(chan string)

	// Launch 3 APIs at the EXACT SAME TIME!
	go fetchPrice("Amazon  (Takes 3s)", 3*time.Second, priceChan)
	go fetchPrice("eBay    (Takes 1s)", 1*time.Second, priceChan)
	go fetchPrice("Walmart (Takes 2s)", 2*time.Second, priceChan)

	// In the Goroutines section, we used time.Sleep(3 * time.Second) to wait for them.
	// That was dangerous! What if Amazon took 4 seconds? Or what if it only took 1 second?
	// By using Channels, we just listen to the pipe 3 times.
	// The program will wait EXACTLY as long as it needs to, and not a millisecond longer!

	msg1 := <-priceChan // Freezes until the FASTEST one (eBay) finishes in 1 sec
	fmt.Println(msg1)

	msg2 := <-priceChan // Freezes until the 2nd fastest one (Walmart) finishes in 2 sec
	fmt.Println(msg2)

	msg3 := <-priceChan // Freezes until the SLOWEST one (Amazon) finishes in 3 sec
	fmt.Println(msg3)

	fmt.Println("All APIs finished! Program exiting instantly.")
}
