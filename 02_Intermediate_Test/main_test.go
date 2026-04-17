package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
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
		var pe PaymentError
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
	if resp.Message != "API is running!" {
		t.Errorf("Expected correct message from API, got: %s", resp.Message)
	}
}

func TestExtractEmails(t *testing.T) {
	text := "Reach us at support@acme.com or sales@acme.com for help."

	emails := ExtractEmails(text)
	if len(emails) != 2 {
		t.Fatalf("Expected 2 emails, got %d: %v", len(emails), emails)
	}
	if emails[0] != "support@acme.com" {
		t.Errorf("Expected first email 'support@acme.com', got %q", emails[0])
	}
	if emails[1] != "sales@acme.com" {
		t.Errorf("Expected second email 'sales@acme.com', got %q", emails[1])
	}

	// No emails in text
	noneFound := ExtractEmails("no emails here")
	if len(noneFound) != 0 {
		t.Errorf("Expected 0 emails from text without addresses, got %d", len(noneFound))
	}
}

func TestRedactEmails(t *testing.T) {
	input := "Contact bob@example.com or alice@example.org for details."
	result := RedactEmails(input)

	expected := "Contact [REDACTED] or [REDACTED] for details."
	if result != expected {
		t.Errorf("RedactEmails returned %q, want %q", result, expected)
	}
}

func TestIsValidPhone(t *testing.T) {
	valid := []string{"212-555-1234", "800-555-9999", "+1 212-555-1234"}
	for _, p := range valid {
		if !IsValidPhone(p) {
			t.Errorf("IsValidPhone(%q) = false, want true", p)
		}
	}

	invalid := []string{"123-456-7890", "not-a-phone", "000-000-0000"}
	for _, p := range invalid {
		if IsValidPhone(p) {
			t.Errorf("IsValidPhone(%q) = true, want false", p)
		}
	}
}

func TestFetchURL(t *testing.T) {
	// Use httptest so tests work offline
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	code, err := FetchURL(ts.URL)
	if err != nil {
		t.Fatalf("FetchURL returned error: %v", err)
	}
	if code != http.StatusOK {
		t.Errorf("Expected 200, got %d", code)
	}
}

func TestGetWithHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != "secret" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := &http.Client{Timeout: 5 * time.Second}

	// Correct header should yield 200
	code, err := GetWithHeader(client, ts.URL, "X-Api-Key", "secret")
	if err != nil {
		t.Fatalf("GetWithHeader returned error: %v", err)
	}
	if code != http.StatusOK {
		t.Errorf("Expected 200 with correct header, got %d", code)
	}

	// Missing / wrong header should yield 401
	code2, _ := GetWithHeader(client, ts.URL, "X-Api-Key", "wrong")
	if code2 != http.StatusUnauthorized {
		t.Errorf("Expected 401 with wrong header, got %d", code2)
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_KEY_PRESENT", "hello")
	defer os.Unsetenv("TEST_KEY_PRESENT")

	val := GetEnvOrDefault("TEST_KEY_PRESENT", "default")
	if val != "hello" {
		t.Errorf("GetEnvOrDefault returned %q, want %q", val, "hello")
	}

	val2 := GetEnvOrDefault("TEST_KEY_ABSENT_XYZ", "fallback")
	if val2 != "fallback" {
		t.Errorf("GetEnvOrDefault returned %q, want %q", val2, "fallback")
	}
}

func TestRequireEnv(t *testing.T) {
	os.Setenv("REQUIRED_KEY_123", "myvalue")
	defer os.Unsetenv("REQUIRED_KEY_123")

	val, err := RequireEnv("REQUIRED_KEY_123")
	if err != nil {
		t.Fatalf("RequireEnv returned error for set key: %v", err)
	}
	if val != "myvalue" {
		t.Errorf("RequireEnv returned %q, want %q", val, "myvalue")
	}

	_, err2 := RequireEnv("DEFINITELY_NOT_SET_KEY_XYZ_999")
	if err2 == nil {
		t.Error("RequireEnv should return an error for an unset key, got nil")
	}
}

