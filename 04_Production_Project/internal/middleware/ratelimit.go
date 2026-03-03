package middleware

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter holds a rate limiter per client IP using the Token Bucket algorithm.
type IPRateLimiter struct {
	ips   map[string]*rate.Limiter
	mu    sync.RWMutex
	rate  rate.Limit
	burst int
}

// NewIPRateLimiter creates a new rate limiter with the given tokens-per-second rate and burst size.
func NewIPRateLimiter(r rate.Limit, burst int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:   make(map[string]*rate.Limiter),
		rate:  r,
		burst: burst,
	}
}

// getLimiter returns the rate limiter for the given IP, creating one if it doesn't exist.
func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]
	if exists {
		i.mu.RUnlock()
		return limiter
	}
	i.mu.RUnlock()

	i.mu.Lock()
	limiter = rate.NewLimiter(i.rate, i.burst)
	i.ips[ip] = limiter
	i.mu.Unlock()
	return limiter
}

// RateLimiter returns a middleware that limits requests per IP address.
func RateLimiter(limiter *IPRateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			lim := limiter.getLimiter(ip)
			if !lim.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
