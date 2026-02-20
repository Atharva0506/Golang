package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func main() {
	num1 := 10.12
	num2 := 20.12

	fmt.Println("Addition: ", num1+num2)
	fmt.Println("Subtraction: ", num1-num2)
	fmt.Println("Multiplication: ", num1*num2)
	fmt.Println("Division: ", num1/num2)
	fmt.Println("Remainder: ", int(num1)%int(num2))

	fmt.Println("Rounding down (Floor)", math.Floor(num1))
	fmt.Println("Rounding up (Ceil)", math.Ceil(num1))
	fmt.Println("Absolute value", math.Abs(num1))
	fmt.Println("Max between two num", math.Max(num1, num2))
	fmt.Println("Min between two num", math.Min(num1, num2))
	fmt.Println("Square root", math.Sqrt(num1))
	fmt.Println("Exponent", math.Pow(num1, num2))

	// rand

	randomInt := rand.Intn(10)
	fmt.Println("Random int: ", randomInt)

	//dice roll
	diceRoll := rand.Intn(6) + 1
	fmt.Println("Dice roll: ", diceRoll)

	//crypto
	secureRandomInt, err := cryptorand.Int(cryptorand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Secure random int: ", secureRandomInt)

	/*
		Difference between rand and crypto

		rand is not secure and is not suitable for cryptographic use
		crypto/rand is secure and is suitable for cryptographic use

		- How they work
			- rand uses a seed value to generate random numbers
			- crypto/rand uses a seed value and a source of entropy to generate random numbers
	*/
}
