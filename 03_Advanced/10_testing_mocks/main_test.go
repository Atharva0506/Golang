package main

import (
	"errors"
	"testing"
)

type MockSender struct{}

func (m *MockSender) Send(msg string) error {
	return errors.New("fake network failure")
}

func TestAlertFailure(t *testing.T) {
	service := NotificationService{Sender: &MockSender{}}

	err := service.Alert("Warning!")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
