package alacpretty

import (
	ansi "github.com/solidiquis/ansigo"
	"os"
	"os/signal"
	"syscall"
)

func ListenForExit(done chan os.Signal) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	done <- sig
}

func Cleanup() {
	ansi.CursorShow()
}
