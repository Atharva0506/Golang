package advanced_test

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// ==========================================
// 1. DOMAIN MODELS & INTERFACES (04_interfaces_advanced)
// ==========================================

// EmailSender Interface allows Mocking (10_testing_mocks)
type EmailSender interface {
	Send(ctx context.Context, to string, body string) error
}

// User Model with Struct Tags for Reflection (05_reflection)
type User struct {
	ID    int    `db:"id"`
	Email string `db:"email"`
	Name  string `db:"name"`
}

// DataStore Interface protects us from real Database crashes! (10_testing_mocks)
type DataStore interface {
	SaveUser(ctx context.Context, query string, args ...any) error
	GetUser(id int) string
}

// ==========================================
// 2. MOCKS (10_testing_mocks)
// ==========================================

// MockEmailSender implements EmailSender for testing without real network calls
type MockEmailSender struct {
	mu           sync.Mutex
	SentMessages []string
}

func (m *MockEmailSender) Send(ctx context.Context, to string, body string) error {
	// (01_context) Respect context cancellation!
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		m.mu.Lock()
		m.SentMessages = append(m.SentMessages, to+":"+body)
		m.mu.Unlock()
		return nil
	}
}

// MockDataStore perfectly isolates the test from C compilers!
type MockDataStore struct {
	SavedUsers map[int]string
}

func (m *MockDataStore) SaveUser(ctx context.Context, query string, args ...any) error {
	id := args[0].(int) // Type Assertion! (04_interfaces_advanced)
	name := args[2].(string)

	if m.SavedUsers == nil {
		m.SavedUsers = make(map[int]string)
	}
	m.SavedUsers[id] = name
	return nil
}

func (m *MockDataStore) GetUser(id int) string {
	return m.SavedUsers[id]
}

// ==========================================
// 3. CORE SERVICE WITH DEPENDENCY INJECTION (06_design_patterns)
// ==========================================

type NotificationService struct {
	DB     DataStore
	Sender EmailSender
}

// NewNotificationService uses Functional Options Pattern (06_design_patterns)
type ServiceOption func(*NotificationService)

func WithMockSender(sender EmailSender) ServiceOption {
	return func(ns *NotificationService) {
		ns.Sender = sender
	}
}

func NewNotificationService(db DataStore, opts ...ServiceOption) *NotificationService {
	ns := &NotificationService{
		DB: db,
	}
	for _, opt := range opts {
		opt(ns)
	}
	return ns
}

// ==========================================
// 4. BUSINESS LOGIC
// ==========================================

// ProcessNewUser saves to DB (09_database_sql) and sends an Email
func (s *NotificationService) ProcessNewUser(ctx context.Context, u *User) error {
	// 4A. Reflection to dynamically generate SQL query (05_reflection)
	t := reflect.TypeOf(*u)
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		columns = append(columns, t.Field(i).Tag.Get("db"))
	}
	// (03_buffers) strings.Builder for fast concatenation
	var sb strings.Builder
	sb.WriteString("INSERT INTO users (")
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(") VALUES (?, ?, ?)") // Native SQLite syntax

	// 4B. Database Transaction execution via Interface (04_interfaces_advanced)
	// Even though we aren't using a real transaction object here for simplicity,
	// In production, `DataStore` would orchestrate the Commit/Rollback.
	err := s.DB.SaveUser(ctx, sb.String(), u.ID, u.Email, u.Name)
	if err != nil {
		return err
	}

	// 4C. Send Email using injected Mock (10_testing_mocks)
	err = s.Sender.Send(ctx, u.Email, "Welcome "+u.Name+"!")
	if err != nil {
		return err
	}

	return nil
}

// ==========================================
// 5. THE MASTER INTEGRATION TEST
// ==========================================

// TestEndToEndIntegration tests the entire Advanced Module flow in a single interconnected pipeline
func TestEndToEndIntegration(t *testing.T) {
	// Step 1: Memory Pointers (02_pointers)
	// We create the user as a pointer to avoid copying struct memory.
	testUser := &User{
		ID:    1,
		Email: "gopher@golang.org",
		Name:  "Gopher",
	}

	// Step 2: Database Setup (10_testing_mocks)
	// Instead of a real DB which requires C compilers on Windows, we use our MockDataStore!
	mockDB := &MockDataStore{}

	// Step 3: Dependency Injection & Mocks (06_design_patterns & 10_testing_mocks)
	mockSender := &MockEmailSender{}
	service := NewNotificationService(mockDB, WithMockSender(mockSender))

	// Step 4: Context Timeouts (01_context)
	// We ensure this test doesn't hang forever!
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Step 5: Execute Core Logic
	// This uses Reflection (05) to build queries, and Transactions (09) safely!
	err := service.ProcessNewUser(ctx, testUser)
	if err != nil {
		t.Fatalf("ProcessNewUser failed: %v", err)
	}

	// Step 6: Verify Database State
	dbName := mockDB.GetUser(testUser.ID)
	if dbName != "Gopher" {
		t.Errorf("Expected name Gopher, got %s", dbName)
	}

	// Step 7: Verify Mock State
	mockSender.mu.Lock()
	defer mockSender.mu.Unlock()
	if len(mockSender.SentMessages) != 1 {
		t.Fatalf("Expected 1 email sent, got %d", len(mockSender.SentMessages))
	}
	expectedEmail := "gopher@golang.org:Welcome Gopher!"
	if mockSender.SentMessages[0] != expectedEmail {
		t.Errorf("Expected email '%s', got '%s'", expectedEmail, mockSender.SentMessages[0])
	}
}

// ==========================================
// 6. BENCHMARKING (11_benchmarking)
// ==========================================

// BenchmarkJSONReflection proves why Reflection (while powerful) is slow.
func BenchmarkJSONReflection(b *testing.B) {
	u := User{ID: 1, Email: "test@test.com", Name: "Test"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// json.Marshal relies heavily on reflect.TypeOf under the hood!
		_, _ = json.Marshal(u)
	}
}
