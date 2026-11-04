package goblox

import (
	"context"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"
)

var (
	baseRetryDelay = 500 * time.Millisecond
	maxRetryDelay  = 30 * time.Second
)

type RateLimit struct {
	Limit     string
	Remaining int
	Reset     time.Duration
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
}

func (p RetryPolicy) backoff(attempt int) time.Duration {
	if attempt < 32 {
		if d := baseRetryDelay << attempt; d > 0 && d < maxRetryDelay {
			maxRetryDelay = d
		}
	}

	return time.Duration(rand.Int64N(int64(maxRetryDelay))) + time.Millisecond
}

// Next reports whether a request should be sent again after an attempt.
// For the first send, the returned duration will be 0.
func (r RetryPolicy) Next(attempt int, req *http.Request, resp *http.Response, err error) (time.Duration, bool) {
	if err != nil || attempt >= r.MaxRetries {
		return 0, false
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		return 0, false
	}

	delay := parseRetryAfter(resp.Header.Get("Retry-After"))
	if delay <= 0 {
		delay = r.backoff(attempt)

		if ratelimit := ParseRateLimit(resp.Header); ratelimit.Remaining == 0 && ratelimit.Reset > 0 {
			delay = ratelimit.Reset
		}
	}
	if delay > maxRetryDelay {
		return 0, false
	}

	return delay, true
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
