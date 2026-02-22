package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 1. Basic Error Return
func login(username, password string) error {
	// Logical Gotcha: It should be OR (||), not AND (&&)!
	// If the username is blank OR the password is wrong, deny them.
	if username == "" || password != "secret123" {
		return errors.New("Invalid credentials")
	}
	return nil
}

func ParseAge(ageStr string) (int, error) {
	num, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func main() {
	// 2. You MUST check the error returned by the function!
	// If you just write `login("xyz", "asas")`, Go silently ignores any errors returned.
	err := login("xyz", "asas")
	if err != nil {
		fmt.Println("Login Failed:", err)
	} else {
		fmt.Println("Login Successful!")
	}

	age, err := ParseAge("12")
	if err != nil {
		fmt.Println("Error while age parse:", err)
		// JOB TIP: You MUST `return` here!
		// If you just print the error and don't exit, the program will keep running below
		// and try to use `age` (which will just be a blank 0)!
		return
	}
	fmt.Printf("Parsed Age: %v\n", age)

}
