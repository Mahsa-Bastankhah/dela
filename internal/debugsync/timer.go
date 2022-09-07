package debugsync

import (
	"runtime/debug"
	"time"

	"go.dedis.ch/dela"
)

const mutexTimeout = 2 * time.Hour

func startLockTimer(msg string) chan struct{} {
	done := make(chan struct{})

	stack := debug.Stack()
	go func() {
		select {
		case <-time.After(mutexTimeout):
			dela.Logger.Error().Msgf("%v : %v", msg, string(stack))
			return
		case <-done:
			return
		}
	}()

	return done
}
