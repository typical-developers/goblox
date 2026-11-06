package goblox

import (
	"context"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"
)

const (
	baseRetryDelay = 500 * time.Millisecond
	maxRetryDelay  = 30 * time.Second
)

type RateLimit struct {
	Limit     string
	Remaining int
	Reset     time.Duration
}

// isIdempotent will determine if the current request is idempotent
// based on the request's method.
func isIdempotent(method string) bool {
	switch method {
	case http.MethodGet,
		http.MethodHead,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodTrace:
		return true
	default:
		return false
	}
}

func ParseRateLimit(h http.Header) RateLimit {
	ratelimit := RateLimit{
		Limit:     h.Get("x-ratelimit-limit"),
		Remaining: -1,
	}

	if v, err := strconv.Atoi(h.Get("x-ratelimit-remaining")); err == nil {
		ratelimit.Remaining = v
	}
	if v, err := strconv.Atoi(h.Get("x-ratelimit-reset")); err == nil {
		ratelimit.Reset = time.Duration(v) * time.Second
	}

	return ratelimit
}

type RetryPolicy struct {
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
	// RetryOnServerError will retry 502, 503, and 504 status codes
	// asusming that the HTTP method is considered idempotent.
	//
	// GET, HEAD, PUT, DELETE, OPTIONS, and TRACE are labeled idempotent.
	//
	// 429s are always retried.
	RetryServerErrors bool

	OnRetry func(attempt int, delay time.Duration, resp *http.Response)
}

var DefaultRetryPolicy = RetryPolicy{
	MaxRetries:        5,
	BaseDelay:         baseRetryDelay,
	MaxDelay:          maxRetryDelay,
	RetryServerErrors: false,
}

func (r RetryPolicy) baseDelay() time.Duration {
	if r.BaseDelay <= 0 {
		return baseRetryDelay
	}

	return r.BaseDelay
}

func (r RetryPolicy) maxDelay() time.Duration {
	if r.MaxDelay <= 0 {
		return maxRetryDelay
	}

	return r.MaxDelay
}

// backoff will calculate a jittered duration based on the current attempt.
func (r RetryPolicy) backoff(attempt int) time.Duration {
	base := float64(r.baseDelay())
	maxDelay := float64(r.maxDelay())

	jitter := rand.Float64()
	ceiling := min(base*math.Pow(2, float64(attempt)), maxDelay)

	return time.Duration(jitter * ceiling)
}

// Next returns the retry delay for a request and if it should be made again.
// The first send, the return duration will always be 0.
func (r RetryPolicy) Next(attempt int, req *http.Request, resp *http.Response, err error) (time.Duration, bool) {
	if err != nil || attempt >= r.MaxRetries {
		return 0, false
	}

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		delay := parseRetryAfter(resp.Header.Get("Retry-After"))
		if delay <= 0 {
			delay = r.backoff(attempt)

			if ratelimit := ParseRateLimit(resp.Header); ratelimit.Remaining == 0 && ratelimit.Reset > 0 {
				delay = ratelimit.Reset
			}
		}

		if delay > r.maxDelay() {
			return 0, false
		}

		return delay, true
	case http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		if !r.RetryServerErrors {
			return 0, false
		}

		if !isIdempotent(req.Method) {
			return 0, false
		}

		delay := r.backoff(attempt)
		if delay > r.maxDelay() {
			return 0, false
		}

		return delay, true
	}

	return 0, false
}

// Sleeps waits for duration or until context is done.
func Sleep(ctx context.Context, d time.Duration) error {
	if deadline, ok := ctx.Deadline(); ok && time.Until(deadline) < d {
		return context.DeadlineExceeded
	}

	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
