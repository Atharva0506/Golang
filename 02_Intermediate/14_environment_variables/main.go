package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

// AppConfig holds all configuration for our application.
// This is the idiomatic Go pattern: read env vars ONCE at startup,
// store them in a typed struct, and pass the struct around.
// Never call os.Getenv() deep inside business logic!
type AppConfig struct {
	DatabaseURL string
	Port        int
	Debug       bool
	APIKey      string
}

// LoadConfig reads all required environment variables into a typed struct.
// Returns an error if a required variable is missing or has an invalid value.
func LoadConfig() (*AppConfig, error) {
	// 1. os.Getenv — returns the value, or "" if the variable is not set.
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://localhost:5432/mydb" // sensible default
	}

	// 2. os.LookupEnv — the safer alternative.
	// It returns (value, true) if the variable is set, or ("", false) if it is not.
	// This lets you distinguish between "not set" and "set to empty string".
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		return nil, fmt.Errorf("required env var API_KEY is not set")
	}

	// 3. Parsing typed values — env vars are always strings, so we must convert them.
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT %q: %w", portStr, err)
	}

	debugStr := os.Getenv("DEBUG")
	debug, _ := strconv.ParseBool(debugStr) // ParseBool handles "true", "1", "false", "0"

	return &AppConfig{
		DatabaseURL: dbURL,
		Port:        port,
		Debug:       debug,
		APIKey:      apiKey,
	}, nil
}

func main() {
	// 4. os.Setenv — set an env var in the current process (useful in tests).
	os.Setenv("API_KEY", "super-secret-key-123")
	os.Setenv("PORT", "9090")
	os.Setenv("DEBUG", "true")

	cfg, err := LoadConfig()
	if err != nil {
		slog.Error("configuration error", "error", err)
		os.Exit(1)
	}

	slog.Info("config loaded",
		"database_url", cfg.DatabaseURL,
		"port", cfg.Port,
		"debug", cfg.Debug,
	)
	// NOTE: Never log sensitive values like API keys in production!
	slog.Info("api key loaded", "key_length", len(cfg.APIKey))

	// 5. os.Environ — get ALL environment variables as "KEY=VALUE" strings.
	fmt.Println("\n--- All ENV vars (first 3) ---")
	for i, kv := range os.Environ() {
		if i >= 3 {
			break
		}
		fmt.Println(kv)
	}

	// 6. os.Unsetenv — remove an env var from the current process.
	os.Unsetenv("API_KEY")
	if _, ok := os.LookupEnv("API_KEY"); !ok {
		fmt.Println("\nAPI_KEY has been unset successfully!")
	}
}
