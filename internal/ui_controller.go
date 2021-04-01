package alacpretty

import (
	us "github.com/solidiquis/alacpretty_v2/internal/uistructs"
	ansi "github.com/solidiquis/ansigo"
	"os"
)

type UIComponent interface {
	Update(string)
	Render()
}

func ListenForInput(config []byte, uicomps ...UIComponent) {
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
			switch b := key[0]; b {
			case KEY_RETURN:
				newThm := FormatTheme(activeComp.(*us.List).Items[activeComp.(*us.List).Selected])
				updatedConf := ChangeTheme(config, newThm)
				ApplyChanges(updatedConf)
			case KEY_ESC:
				key = SpecialKeys(stdin)
				activeComp.(*us.List).Update(key)
			default:
				activeComp.(*us.List).Update(key)
			}
		case <-done:
			Cleanup()
			break event
		}
	}
}

func SpecialKeys(stdin chan string) string {
	secKey := <-stdin
	if secKey[0] == 91 {
		switch tertKey := <-stdin; tertKey[0] {
		case 66:
			return "<Down>"
		case 65:
			return "<Up>"
		case 67:
			return "<Right>"
		case 68:
			return "<Left>"
		}
	}

	return ""
}
