package main

import (
	"errors"
	"fmt"
)

type PaymentError struct {
	ErrorCode int
	Reason    string
}

// 1. Fulfil the `error` interface by adding an Error() method
func (p PaymentError) Error() string {
	return fmt.Sprintf("Payment Failed! Code: %d | Message: %s", p.ErrorCode, p.Reason)
}

func ChargeCustomer(balance int) error {
	if balance < 50 {
		// 2. Return our custom struct instead of errors.New!
		return PaymentError{ErrorCode: 402, Reason: "Balance is less than 50"}
	}
	return nil
}
func main() {
	err := ChargeCustomer(49)
	if err != nil {
		var paymentError PaymentError

		// 3. Extracting the rich custom error using errors.As
		if errors.As(err, &paymentError) {
			// Because we unpacked it, we can access `.ErrorCode` directly!
			fmt.Printf("Payment Declined! Exact HTTP Error Code to send to frontend: %d\n", paymentError.ErrorCode)
			fmt.Printf("Reason to show user: %s\n", paymentError.Reason)
		} else {
			fmt.Println("Generic Server Error:", err)
		}

		return // Job Tip: ALWAYS return after an error so it doesn't continue down!
	}

	fmt.Println("Payment Successful!")
}
