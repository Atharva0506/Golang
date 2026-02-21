package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func checkAgeGroup(age int) string {
	if age <= 0 {
		return "Invlaid Age"
	} else if age < 13 {
		return "Child"
	} else if age >= 13 && age <= 19 {
		return "Teen"
	} else {
		return "Adult"
	}

}
func main() {
	fmt.Print(checkAgeGroup(20))

	// 2. Initialization Statement (crucial for Go jobs)
	// We create `num` and `err`, and IMMEDIATELY check if `err` exists.
	if num, err := rand.Int(rand.Reader, big.NewInt(5)); err != nil {
		fmt.Println("Error while generating num", err)
	} else if num.Int64() < 3 {
		// Because num is a *big.Int, we must convert it to a regular Int64 to compare it to a normal number like 3
		fmt.Println("\nNumber is less than 3: ", num)
	} else {
		fmt.Println("\nNumber is greater than or equal to 3:", num)
	}

}
