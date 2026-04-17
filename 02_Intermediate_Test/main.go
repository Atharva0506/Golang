package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

// Task 1: Interfaces and Custom Errors
// 1. Create a Custom Error strictly named 'PaymentError' with two fields: 'Code' (int) and 'Message' (string).
// 2. Add an 'Error() string' method to it so it fulfills the error interface. It should return a formatted string.
// 3. Create an interface named 'PaymentProcessor' with a single method: `Pay(amount float64) error`.
// 4. Create a struct named 'StripeProcessor'.
// 5. Add the 'Pay' method to 'StripeProcessor' so it fulfills the interface.
//    -> Inside Pay: If the amount is less than 10.0, return a 'PaymentError' with Code: 400, Message: "Minimum payment is $10".
//    -> Otherwise, return nil.

// Task 2: Concurrency (Goroutines, WaitGroups, Mutex)
// 1. Create a function: 'ProcessBatch(amounts []float64) float64'.
// 2. Inside, use a sync.WaitGroup and a sync.Mutex.
// 3. Create a 'total' float64 variable starting at 0.
// 4. For every amount in the 'amounts' slice, launch an ANONYMOUS Goroutine.
// 5. Inside the Goroutine, if the amount >= 10.0, safely add it to 'total' using the Mutex.
// 6. Make sure to call wg.Add(1) before launching, and defer wg.Done() inside the Goroutine.
// 7. Call wg.Wait() to block until all workers finish, then return 'total'.

// Task 3: Generics and JSON
// 1. Create a function: 'FilterData[T any](jsonData []byte, filterFunc func(item T) bool) ([]T, error)'.
// 2. Unmarshal the 'jsonData' into a slice of type 'T' (e.g. `var items []T`).
// 3. If unmarshaling fails, return nil and the error.
// 4. If it succeeds, loop through 'items'. Pass each item to 'filterFunc'.
// 5. If 'filterFunc' returns true, append the item to a new slice called 'filtered'.
// 6. Return 'filtered' and nil.

// Task 4: Building an API
// 1. Create a struct named 'StatusResponse' with a boolean field 'Success' (JSON tag: "success") and string field 'Message' (JSON tag: "message").
// 2. Create a handler function: 'StatusHandler(w http.ResponseWriter, r *http.Request)'.
// 3. Inside the handler, set the Content-Type header to 'application/json'.
// 4. Create an instance of 'StatusResponse' (Success: true, Message: "API is running!").
// 5. Use json.NewEncoder to encode and send the response.
// 6. Create a function 'SetupRouter() *http.ServeMux'. Inside to create a new Mux (`mux := http.NewServeMux()`)
//    and register the route "GET /status" to your StatusHandler. Return the `mux`!

// Task 5: Regular Expressions (12_regex)
// 1. Create a function 'ExtractEmails(text string) []string'.
//    - Compile a regular expression that matches email addresses.
//    - Return ALL matching email addresses found in 'text' using FindAllString.
//    - Return nil if no emails are found.
//
// 2. Create a function 'RedactEmails(text string) string'.
//    - Replace every email address in 'text' with the literal string "[REDACTED]".
//    - Return the redacted string.
//
// 3. Create a function 'IsValidPhone(phone string) bool'.
//    - Return true if 'phone' matches the pattern: optional "+1 ", then 10 digits
//      formatted as NXX-NXX-XXXX where N is 2-9.
//    - Use the pattern: `^\+?1?\s?\(?[2-9]\d{2}\)?[\s\-]?[2-9]\d{2}\-\d{4}$`

func ExtractEmails(text string) []string {
	// TODO: Implement
	return nil
}

func RedactEmails(text string) string {
	// TODO: Implement
	return text
}

func IsValidPhone(phone string) bool {
	// TODO: Implement
	return false
}

// Task 6: HTTP Client (13_http_client)
// 1. Create a function 'FetchURL(url string) (int, error)'.
//    - Create an *http.Client with a 5-second timeout.
//    - Perform a GET request to 'url'.
//    - Return the HTTP status code and any error.
//    - IMPORTANT: Always close resp.Body!
//
// 2. Create a function 'GetWithHeader(client *http.Client, url, headerKey, headerValue string) (int, error)'.
//    - Build an *http.Request using http.NewRequest.
//    - Set the custom header on the request.
//    - Execute it with the provided client.
//    - Return the status code and any error.

func FetchURL(url string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func GetWithHeader(client *http.Client, url, headerKey, headerValue string) (int, error) {
	// TODO: Implement
	return 0, nil
}

// Task 7: Environment Variables (14_environment_variables)
// 1. Create a function 'GetEnvOrDefault(key, defaultValue string) string'.
//    - Return the environment variable named 'key' if it is set and non-empty.
//    - Otherwise return 'defaultValue'.
//
// 2. Create a function 'RequireEnv(key string) (string, error)'.
//    - Use os.LookupEnv to check if 'key' is defined.
//    - If it is NOT defined, return an error: fmt.Errorf("required env var %q is not set", key).
//    - If it IS defined, return the value.

func GetEnvOrDefault(key, defaultValue string) string {
	// TODO: Implement
	return defaultValue
}

func RequireEnv(key string) (string, error) {
	// TODO: Implement
	return "", nil
}

// suppress unused import errors until students implement the tasks
var _ = regexp.MustCompile
var _ = os.Getenv

func main() {
	fmt.Println("Good luck! Run `go test -v` in this directory to check your answers.")
}
