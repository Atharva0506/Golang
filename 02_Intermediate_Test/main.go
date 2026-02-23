package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Task 1: Interfaces and Custom Errors
//  1. Create a Custom Error strictly named 'PaymentError' with two fields: 'Code' (int) and 'Message' (string).
//  2. Add an 'Error() string' method to it so it fulfills the error interface. It should return a formatted string.
//  3. Create an interface named 'PaymentProcessor' with a single method: `Pay(amount float64) error`.
//  4. Create a struct named 'StripeProcessor'.
//  5. Add the 'Pay' method to 'StripeProcessor' so it fulfills the interface.
//     -> Inside Pay: If the amount is less than 10.0, return a 'PaymentError' with Code: 400, Message: "Minimum payment is $10".
//     -> Otherwise, return nil.

type PaymentError struct {
	Code    int
	Message string
}

type StripeProcessor struct {
}
type PaymentProcessor interface {
	Pay(amount float64) error
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("Payment Error Code: %d | Message : %s", e.Code, e.Message)

}
func (s StripeProcessor) Pay(amount float64) error {
	if amount < 10.0 {
		return &PaymentError{
			Code:    400,
			Message: "Minimum payment is $10",
		}
	}
	return nil
}

// Task 2: Concurrency (Goroutines, WaitGroups, Mutex)
// 1. Create a function: 'ProcessBatch(amounts []float64) float64'.
// 2. Inside, use a sync.WaitGroup and a sync.Mutex.
// 3. Create a 'total' float64 variable starting at 0.
// 4. For every amount in the 'amounts' slice, launch an ANONYMOUS Goroutine.
// 5. Inside the Goroutine, if the amount >= 10.0, safely add it to 'total' using the Mutex.
// 6. Make sure to call wg.Add(1) before launching, and defer wg.Done() inside the Goroutine.
// 7. Call wg.Wait() to block until all workers finish, then return 'total'.

func ProcessBatch(amounts []float64) float64 {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var total float64
	for _, amt := range amounts {
		wg.Add(1)
		go func(a float64) {
			defer wg.Done()
			if a >= 10.0 {
				mu.Lock()
				total += a
				mu.Unlock()
			}
		}(amt)

	}
	wg.Wait()
	return total
}

// Task 3: Generics and JSON
// 1. Create a function: 'FilterData[T any](jsonData []byte, filterFunc func(item T) bool) ([]T, error)'.
// 2. Unmarshal the 'jsonData' into a slice of type 'T' (e.g. `var items []T`).
// 3. If unmarshaling fails, return nil and the error.
// 4. If it succeeds, loop through 'items'. Pass each item to 'filterFunc'.
// 5. If 'filterFunc' returns true, append the item to a new slice called 'filtered'.
// 6. Return 'filtered' and nil.
func FilterData[T any](jsonData []byte, filterFunc func(item T) bool) ([]T, error) {
	var items []T
	filtered := make([]T, 0)

	err := json.Unmarshal(jsonData, &items)
	if err != nil {
		return nil, err
	}
	for _, data := range items {
		if filterFunc(data) {
			filtered = append(filtered, data)
		}
	}
	return filtered, nil
}

// Task 4: Building an API
//  1. Create a struct named 'StatusResponse' with a boolean field 'Success' (JSON tag: "success") and string field 'Message' (JSON tag: "message").
//  2. Create a handler function: 'StatusHandler(w http.ResponseWriter, r *http.Request)'.
//  3. Inside the handler, set the Content-Type header to 'application/json'.
//  4. Create an instance of 'StatusResponse' (Success: true, Message: "API is running!").
//  5. Use json.NewEncoder to encode and send the response.
//  6. Create a function 'SetupRouter() *http.ServeMux'. Inside to create a new Mux (`mux := http.NewServeMux()`)
//     and register the route "GET /status" to your StatusHandler. Return the `mux`!
type StatusResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(StatusResponse{Success: true, Message: "API Is running!"})

}
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /status", StatusHandler)
	return mux
}
func main() {
	fmt.Println("Good luck! Run `go test -v` in this directory to check your answers.")
}
