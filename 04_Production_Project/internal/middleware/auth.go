package middleware

// TODO: JWT Authentication Middleware
// Responsibilities:
// - Intercept HTTP or gRPC requests before they reach the Handler
// - Verify Authorization headers (Bearer tokens) using internal pkg/auth tools
// - Parse claims and inject the authenticated UserID into the context.Context
// - Reject unauthorized connections instantly
