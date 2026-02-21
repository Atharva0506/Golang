package main

import "fmt"

type Account struct {
	name    string
	balance int
	id      int
}

// 1. Value Receiver Method (Read-only)
func (u Account) GetInfo() string {
	// Job Tip: Never use string(rune(int)) to convert a number to a string!
	// That tries to find the unicode character for that number.
	// Instead, use fmt.Sprintf() for super clean string formatting:
	return fmt.Sprintf("Name: %s | Balance: $%d | ID: %d", u.name, u.balance, u.id)
}

// 2. Pointer Receiver Method (Modifies the struct)
func (u *Account) DepositCash(amt int) {
	u.balance += amt
}
func main() {
	luffy := Account{
		name:    "Monky D. Luffy",
		balance: 20000,
		id:      1,
	}

	fmt.Println("Before Deposit:")
	fmt.Println(luffy.GetInfo())

	luffy.DepositCash(100000)

	fmt.Println("\nAfter Deposit:")
	fmt.Println(luffy.GetInfo())
}
