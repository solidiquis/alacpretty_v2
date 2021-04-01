package main

import (
	ap "github.com/solidiquis/alacpretty_v2"
	//ansi "github.com/solidiquis/ansigo"
)

func main() {
	//config := ap.ConfigToBytes()
	//theme, err := ap.CurrentTheme(config)
	//if err != nil {
	//theme = "ARGONAUT"
	//}

	//ap.ChangeTheme(config, theme)

	themes := make([]string, len(ap.Themes))
	i := 0
	for k, _ := range ap.Themes {
		themes[i] = ap.ThemePresenter(k)
		i++
	}

	themesLi := ap.InitList(themes, 0, ap.LIST_FRAME_LENGTH)
	themesLi.Render()

	ap.ListenForInput(themesLi)
}
