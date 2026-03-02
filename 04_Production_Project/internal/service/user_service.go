package service

import (
	"context"
	"strings"
	"time"

	"github.com/Atharva0506/trading_bot/internal/models"
	"github.com/Atharva0506/trading_bot/internal/repository"
	"github.com/Atharva0506/trading_bot/pkg/apperrors"
	"github.com/Atharva0506/trading_bot/pkg/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// TokenPair holds both the access and refresh tokens returned on login.
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserService struct {
	repo          repository.UserRepository
	jwtSecret     string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewUserService(repo repository.UserRepository, jwtSecret string, accessExpiry, refreshExpiry time.Duration) *UserService {
	return &UserService{
		repo:          repo,
		jwtSecret:     jwtSecret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, email, password string) (*models.User, error) {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return nil, apperrors.NewBadRequest("email and password are required")
	}

	// Step 1: Check if the DB lookup itself failed (connection error, timeout, etc.)
	existingUser, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to check existing user")
	}

	// Step 2: If no error and a user was found, it's a conflict.
	if existingUser != nil {
		return nil, apperrors.NewConflict("user already exists")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to hash password")
	}
	user := models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hashPassword),
		Role:         "trader",
		Status:       models.StatusActive,
		CreatedAt:    time.Now(),
	}
	err = s.repo.Create(ctx, &user)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to create user")
	}
	return &user, nil
}

func (s *UserService) LoginUser(ctx context.Context, email, password string) (*TokenPair, error) {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return nil, apperrors.NewBadRequest("email and password are required")
	}

	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, apperrors.NewInternal(err, "server error")
	}
	if user == nil {
		// Use a generic message so attackers can't enumerate valid emails.
		return nil, apperrors.NewUnauthorized("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		// Same generic message — don't reveal which field was wrong.
		return nil, apperrors.NewUnauthorized("invalid email or password")
	}

	// Generate a short-lived access token for API requests.
	accessToken, err := auth.GenerateToken(user.ID, user.Role, s.jwtSecret, "access", s.accessExpiry)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to generate access token")
	}

	// Generate a longer-lived refresh token for obtaining new access tokens.
	refreshToken, err := auth.GenerateToken(user.ID, user.Role, s.jwtSecret, "refresh", s.refreshExpiry)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to generate refresh token")
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
