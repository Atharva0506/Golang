package mocks

import (
	"context"

	"github.com/Atharva0506/trading_bot/internal/models"
	"github.com/google/uuid"
)

// MockUserRepo is a fake implementation of repository.UserRepository for testing.
type MockUserRepo struct {
	Users map[string]*models.User
	Err   error
}

// NewMockUserRepo returns a new MockUserRepo with an initialized map.
func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{
		Users: make(map[string]*models.User),
	}
}

// Create stores a user in the in-memory map.
func (m *MockUserRepo) Create(ctx context.Context, user *models.User) error {
	if m.Err != nil {
		return m.Err
	}
	m.Users[user.Email] = user
	return nil
}

// GetByID searches for a user by UUID.
func (m *MockUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	for _, u := range m.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

// GetByEmail looks up a user by email.
func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	user, exists := m.Users[email]
	if exists {
		return user, nil
	}
	return nil, nil
}
