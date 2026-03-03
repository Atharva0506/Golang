package tests

import (
	"context"
	"testing"
	"time"

	"github.com/Atharva0506/trading_bot/internal/service"
	"github.com/Atharva0506/trading_bot/tests/mocks"
)

// helper to create a fresh mock + service for each test.
func setupUserService() (*mocks.MockUserRepo, *service.UserService) {
	mockRepo := mocks.NewMockUserRepo()
	svc := service.NewUserService(mockRepo, "test-secret", 15*time.Minute, 24*time.Hour)
	return mockRepo, svc
}

// --- Register Tests ---

func TestRegisterUser_Success(t *testing.T) {
	_, svc := setupUserService()

	user, err := svc.RegisterUser(context.Background(), "test@test.com", "password123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user == nil {
		t.Fatal("expected user, got nil")
	}
	if user.Email != "test@test.com" {
		t.Errorf("expected email test@test.com, got %s", user.Email)
	}
	if user.Role != "trader" {
		t.Errorf("expected role trader, got %s", user.Role)
	}
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
	_, svc := setupUserService()

	// Register the first user.
	_, err := svc.RegisterUser(context.Background(), "test@test.com", "password123")
	if err != nil {
		t.Fatalf("first registration failed: %v", err)
	}

	// Try to register again with the same email.
	_, err = svc.RegisterUser(context.Background(), "test@test.com", "password456")
	if err == nil {
		t.Fatal("expected conflict error for duplicate email, got nil")
	}
}

func TestRegisterUser_EmptyEmail(t *testing.T) {
	_, svc := setupUserService()

	_, err := svc.RegisterUser(context.Background(), "", "password123")
	if err == nil {
		t.Fatal("expected error for empty email, got nil")
	}
}

func TestRegisterUser_EmptyPassword(t *testing.T) {
	_, svc := setupUserService()

	_, err := svc.RegisterUser(context.Background(), "test@test.com", "")
	if err == nil {
		t.Fatal("expected error for empty password, got nil")
	}
}

// --- Login Tests ---

func TestLoginUser_Success(t *testing.T) {
	_, svc := setupUserService()

	// Register a user first so we have valid credentials in the mock DB.
	_, err := svc.RegisterUser(context.Background(), "test@test.com", "password123")
	if err != nil {
		t.Fatalf("registration failed: %v", err)
	}

	// Login with correct credentials.
	tokens, err := svc.LoginUser(context.Background(), "test@test.com", "password123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tokens == nil {
		t.Fatal("expected token pair, got nil")
	}
	if tokens.AccessToken == "" {
		t.Error("expected access token, got empty string")
	}
	if tokens.RefreshToken == "" {
		t.Error("expected refresh token, got empty string")
	}
}

func TestLoginUser_WrongPassword(t *testing.T) {
	_, svc := setupUserService()

	_, err := svc.RegisterUser(context.Background(), "test@test.com", "password123")
	if err != nil {
		t.Fatalf("registration failed: %v", err)
	}

	_, err = svc.LoginUser(context.Background(), "test@test.com", "wrongpassword")
	if err == nil {
		t.Fatal("expected error for wrong password, got nil")
	}
}

func TestLoginUser_UserNotFound(t *testing.T) {
	_, svc := setupUserService()

	// Don't register any user — try to login directly.
	_, err := svc.LoginUser(context.Background(), "nobody@test.com", "password123")
	if err == nil {
		t.Fatal("expected error for non-existent user, got nil")
	}
}

func TestLoginUser_EmptyFields(t *testing.T) {
	_, svc := setupUserService()

	_, err := svc.LoginUser(context.Background(), "", "password123")
	if err == nil {
		t.Fatal("expected error for empty email, got nil")
	}

	_, err = svc.LoginUser(context.Background(), "test@test.com", "")
	if err == nil {
		t.Fatal("expected error for empty password, got nil")
	}
}
