package shutdown

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrTermSig = errors.New("termination signal caught")
)

type SignalTrap chan os.Signal

// TermSignalTrap ...
func TermSignalTrap() SignalTrap {
	trap := SignalTrap(make(chan os.Signal, 1))
	signal.Notify(trap, syscall.SIGINT, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGSEGV)

	return trap
}

// Wait ...
func (t SignalTrap) Wait(ctx context.Context) error {
	select {
	case <-t:
		return ErrTermSig
	case <-ctx.Done():
		return ctx.Err()
	}
}
