package main

import "fmt"

func main() {

	// 1. Slice Optimization (Pre-allocation)
	// By using make() with a capacity of 500, we ask the OS for one large block of RAM up front.
	// This prevents Go from having to pause the loop multiple times to reallocate memory as it grows!
	users := make([]int, 0, 500)
	for i := 0; i < 500; i++ {
		users = append(users, i)
	}

	fmt.Println("Length of users slice:", len(users))
	fmt.Println("Capacity of users slice:", cap(users))

	// 2. Buffered Channels
	// A standard channel blocks instantly. By giving it a capacity of 2,
	// we can instantly queue up 2 jobs in memory without needing a worker to immediately read them.
	jobs := make(chan int, 2)

	jobs <- 100
	jobs <- 200
	fmt.Println("Successfully loaded 2 jobs into the buffered channel without deadlocking!")

}
