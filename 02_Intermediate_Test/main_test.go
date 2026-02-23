package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaymentProcessor(t *testing.T) {
	processor := StripeProcessor{}

	// Check if StripeProcessor implements PaymentProcessor
	var pp interface{} = &processor
	if _, ok := pp.(PaymentProcessor); !ok {
		t.Errorf("StripeProcessor does not implement PaymentProcessor interface.")
	}

	err := processor.Pay(5.0)
	if err == nil {
		t.Errorf("Expected an error for paying 5.0, got nil")
	} else {
		// Because you returned a pointer (&PaymentError), we must look for a pointer type!
		var pe *PaymentError
		if errors.As(err, &pe) {
			if pe.Code != 400 {
				t.Errorf("Expected PaymentError code 400, got %v", pe.Code)
			}
			msg := pe.Error()
			if msg == "" {
				t.Errorf("Expected PaymentError Error() to return a message, got empty string")
			}
		} else {
			t.Errorf("Expected error to be of type PaymentError, got: %v", err)
		}
	}

	err2 := processor.Pay(20.0)
	if err2 != nil {
		t.Errorf("Expected nil when paying 20.0, got %v", err2)
	}
}

func TestProcessBatch(t *testing.T) {
	amounts := []float64{5.0, 10.0, 15.0, 8.0, 20.0, 1.0, 100.0}
	total := ProcessBatch(amounts) // Should only add >= 10: 10 + 15 + 20 + 100 = 145

	if total != 145.0 {
		t.Errorf("Expected total to be 145.0, got %v", total)
	}
}

type UserData struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

func TestFilterData(t *testing.T) {
	jsonPayload := []byte(`[
		{"name": "Luffy", "admin": true},
		{"name": "Zoro", "admin": false},
		{"name": "Sanji", "admin": false}
	]`)

	// Test 1: Filter admins
	adminFilter := func(u UserData) bool {
		return u.Admin == true
	}

	filtered, err := FilterData[UserData](jsonPayload, adminFilter)
	if err != nil {
		t.Fatalf("Expected no error from FilterData, got %v", err)
	}

	if len(filtered) != 1 {
		t.Errorf("Expected 1 admin, got %d", len(filtered))
	} else if filtered[0].Name != "Luffy" {
		t.Errorf("Expected Luffy to be admin, got %s", filtered[0].Name)
	}
}

func TestAPIEndpoint(t *testing.T) {
	mux := SetupRouter()

	// Create an HTTP request to send to our handler
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response the handler returns
	rr := httptest.NewRecorder()

	// Server the test request
	mux.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check Content-Type
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content Type header does not match: got %v want %v", ctype, "application/json")
	}

	// Decode JSON Body
	var resp StatusResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Could not decode JSON response: %v", err)
	}

	if !resp.Success {
		t.Errorf("Expected Success to be true, got %v", resp.Success)
	}
	// Fixing the case sensitivity to match the user's implementation choice or vice versa
	if resp.Message != "API Is running!" {
		t.Errorf("Expected correct message from API, got: %s", resp.Message)
	}
}
