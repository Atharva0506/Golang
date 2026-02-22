package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	urls := []string{"google.com", "amazon.com", "github.com"}

	for _, url := range urls {
		// What does wg.Add(1) do?
		// It literally just adds `1` to an internal integer counter inside the `wg` struct.
		// We are saying: "Hey WaitGroup, I am about to launch 1 new worker. Add 1 to your clipboard tally."
		// You MUST call Add() BEFORE you write the `go func()` line!
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fmt.Println("Downloading from: " + u)
			time.Sleep(1 * time.Second)
			fmt.Println("Finished Download: " + u)
		}(url) // We pass `url` into the closure so it guarantees it uses the correct url string!
	}

	// wg.Wait() completely freezes the `main` function right here.
	// It constantly checks the `wg` tally. When `wg.Done()` eventually lowers the tally back down to 0,
	// Wait() unfreezes and allows the program to continue and exit.
	wg.Wait()

	fmt.Println("\nAll downloads completed successfully! Exiting.")
}
