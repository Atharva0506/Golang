package http

// TODO: HTTP Router registration
// Responsibilities:
// - Group application routes by domain (e.g., /api/v1/users)
// - Attach specific Middleware to specific groups (e.g., Auth required for /me)
// - Initialize lightweight multiplexer (like Chi or standard `net/http`)
