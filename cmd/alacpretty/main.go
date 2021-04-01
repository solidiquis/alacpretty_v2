package main

import (
	ap "github.com/solidiquis/alacpretty_v2/internal"
	//ansi "github.com/solidiquis/ansigo"
)

func main() {
	config := ap.ConfigToBytes()

	themes := make([]string, len(ap.Themes))
	i := 0
	for k, _ := range ap.Themes {
		themes[i] = ap.ThemePresenter(k)
		i++
	}

	themesLi := ap.InitList(themes, 0)
	themesLi.Render()

	ap.ListenForInput(config, themesLi)
}
