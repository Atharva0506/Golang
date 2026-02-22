package main

import (
	"fmt"
	"sync"
)

var tickets = 500
var wg sync.WaitGroup
var mu sync.Mutex

func BuyTicket() {
	// Job Tip: defer statements run in LIFO order (Last In, First Out).
	// We defer BOTH `Done` and `Unlock` so that even if the function crashes midway,
	// the lock is released, and the WaitGroup tally goes down!
	defer wg.Done()

	// Lock the door! No other goroutine can pass this line until we call Unlock()
	mu.Lock()
	defer mu.Unlock() // Will automatically unlock right as the function exits

	if tickets > 0 {
		tickets-- // Dangerous memory operation safely protected by padlock
	}
}
func main() {
	// 2. Launch 2000 customers at the exact same time
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go BuyTicket()
	}

	// 3. Wait for all 2000 customers to finish buying
	// BUG FIX: You accidentally put wg.Wait() INSIDE the loop!
	// That meant customer 2 couldn't even spawn until customer 1 finished.
	wg.Wait()
	fmt.Println("ALl ticket sold out: ", tickets)
}
