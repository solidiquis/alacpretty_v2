package alacpretty

import (
	ansi "github.com/solidiquis/ansigo"
	"os"
	"os/signal"
	"syscall"
)

// For graceful exiting of program.
func ListenForExit(done chan os.Signal) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	done <- sig
}

// Restores terminal defaults.
func Cleanup() {
	ansi.CursorShow()
}
