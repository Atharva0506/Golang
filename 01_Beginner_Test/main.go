package main

import (
	"fmt"
)

// Task 1: Structs and Methods
// Create a struct named 'Player' with exportable fields: Name (string), Health (int), Score (int), and IsAlive (bool)

// Write a Pointer Receiver method called 'TakeDamage' for the Player that takes damage (int).
// It should reduce Health by the damage amount.
// If Health drops to 0 or below, set Health to 0 and set IsAlive to false.

// Task 2: Pointers and Math/Crypto
// Write a function called 'HealPlayer' that takes a pointer to a Player struct and an integer for maxHealAmount.
// Inside, generate a random number between 1 and maxHealAmount using `math/rand`.
// Add the random amount to the player's health. Do not let health exceed 100.

// Task 3: Maps and Slices
// Write a function called 'ClassifyScores' that takes a slice of integers (scores).
// It should return a map[string]int where the key is the category ("Low", "Medium", "High") and the value is the count of how many scores fit that category.
// < 50 = "Low", 50-79 = "Medium", 80+ = "High"

func main() {
	fmt.Println("Run `go test` in this directory to see if your code passes!")
}
