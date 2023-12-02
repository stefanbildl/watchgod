package watchgod

import (
	"context"
	"time"
)


type ResetFn func ()
// Extends the parent context with a watchdog timer that can be reset
// using the provided reset function.
// When the watchdog timer with the given duration elapses without being reset in time, 
// the returned context is cancelled
func WithWatchdog(parent context.Context, timeout time.Duration) (context.Context, ResetFn) {
	c, cancel := context.WithCancel(parent)
	reset := make(chan struct{}, 1)
	go func() {
		t := time.NewTimer(timeout)
		b: for {
			select {
				case <- c.Done():
					break b
				case <- t.C:
					cancel()
					break b
				case <- reset:
					if !t.Stop() {
						 <- t.C
					}
					t.Reset(timeout)
			}
		}
	}()
	return c, func() {
		reset <- struct{}{}
	}
}
