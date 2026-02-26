package main

import (
	"errors"
	"testing"
)

// 1. THE MOCK OBJECT (A completely fake struct used only for testing)
type MockSender struct{}

// 2. We attach the exact signature needed to fulfill the `MessageSender` Interface.
// Instead of making real network calls (which costs money and time), we instantly return an error!
func (m *MockSender) Send(msg string) error {
	return errors.New("fake network failure")
}

func TestAlertFailure(t *testing.T) {
	// 3. We inject our fake `MockSender` into the real `NotificationService`
	service := NotificationService{Sender: &MockSender{}}

	// 4. We execute the real method. Under the hood, it triggers our fake Send()!
	err := service.Alert("Warning!")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
