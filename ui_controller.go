package alacpretty

import (
	ansi "github.com/solidiquis/ansigo"
	"os"
)

type UIComponent interface {
	Update(string)
	Render()
}

func ListenForInput(uicomps ...UIComponent) {
	done := make(chan os.Signal, 1)
	go ListenForExit(done)

	activeComp := uicomps[0]
	stdin := make(chan string, 1)

	ansi.CursorHide()
	go ansi.GetChar(stdin)

event:
	for {
		select {
		case key := <-stdin:
			activeComp.(*List).Update(key)
		case <-done:
			Cleanup()
			break event
		}
	}
}
