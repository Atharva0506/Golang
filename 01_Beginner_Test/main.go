package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Task 1: Structs and Methods
// Create a struct named 'Player' with exportable fields: Name (string), Health (int), Score (int), and IsAlive (bool)
type Player struct {
	Name    string
	Health  int
	Score   int
	IsAlive bool
}

// Write a Pointer Receiver method called 'TakeDamage' for the Player that takes damage (int).
// It should reduce Health by the damage amount.
// If Health drops to 0 or below, set Health to 0 and set IsAlive to false.
func (p *Player) TakeDamage(damage int) {
	p.Health -= damage
	if p.Health <= 0 {
		p.Health = 0
		p.IsAlive = false
	}
}

// Task 2: Pointers and Math/Crypto
// Write a function called 'HealPlayer' that takes a pointer to a Player struct and an integer for maxHealAmount.
// Inside, generate a random number between 1 and maxHealAmount using `math/rand`.
// Add the random amount to the player's health. Do not let health exceed 100.
func HealPlayer(p *Player, maxHealAmount int) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(maxHealAmount)))
	if err != nil {
		fmt.Println("Error While Generating number", err)
	}
	num := randomNum.Int64() + 1
	p.Health += int(num)

	if p.Health > 100 {
		p.Health = 100
	}
}

// Task 3: Maps and Slices
// Write a function called 'ClassifyScores' that takes a slice of integers (scores).
// It should return a map[string]int where the key is the category ("Low", "Medium", "High") and the value is the count of how many scores fit that category.
// < 50 = "Low", 50-79 = "Medium", 80+ = "High"
func ClassifyScores(scores []int) map[string]int {
	low := 0
	medium := 0
	high := 0
	for _, s := range scores {
		if s < 50 {
			low++
		} else if s < 79 && s > 50 {
			medium++
		} else {
			high++
		}
	}
	return map[string]int{
		"Low":    low,
		"Medium": medium,
		"High":   high,
	}

}

func ClassifyScoresUsingMap(scores []int) map[string]int {
	counts := make(map[string]int) // Make an empty map
	for _, s := range scores {
		if s < 50 {
			counts["Low"]++ // If "Low" doesn't exist, it starts at 0 automatically!
		} else if s <= 79 { // Using <= ensures we don't accidentally miss exactly 79
			counts["Medium"]++
		} else {
			counts["High"]++
		}
	}

	// Returning the map that already has all your counts compiled instantly!
	return counts
}
func main() {
	fmt.Println("Run `go test` in this directory to see if your code passes!")
}
