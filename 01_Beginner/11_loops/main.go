package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	health := 13

	// "While" style loop
	for health > 0 {
		fmt.Println("Health is:", health)
		health /= 2 // Cuts in half
	}
	fmt.Println("Health Reached Zero")

	// Infinite loop for a guessing game
	secretNumBig, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Printf("Error in Number: %v\n", err)
	}
	// Convert *big.Int to a normal int64 for safe mathematical comparison
	secretNum := secretNumBig.Int64()

	var inp int64
	for {
		fmt.Println("Enter Your Guess (0-99):")
		fmt.Scan(&inp)

		if secretNum == inp {
			fmt.Println("You Won!!!")
			break
		}

		if secretNum < inp {
			fmt.Println("Think Lower")
		} else {
			fmt.Println("Think Higher")
		}
	}

	nums := []int{10, 20, 30}

	for _, n := range nums {
		fmt.Print(" ", n)
	}
}
