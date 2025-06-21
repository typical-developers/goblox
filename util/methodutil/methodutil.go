package methodutil

import (
	"time"
)

// PollMethod is a utility function to help repeatidly call a method until a desired condition is met.
//
// The method will be called at a specified interval. The data returned should be a pointer.
// The handler will return a function method to be called when the desired condition is met.
//
// If the interval is nil, it will default to 5 seconds.
func PollMethod[T any](method func() *T, handler func(*T, func()), interval *time.Duration) {
	if interval == nil {
		i := 5 * time.Second
		interval = &i
	}

	finished := make(chan bool)
	done := func() {
		close(finished)
	}

	defer func() {
		for {
			select {
			case <-finished:

				return
			default:
				d := method()

				handler(d, done)
				time.Sleep(*interval)
			}
		}
	}()
}
