// For manual testing of ui components in isolation.
package main

import (
	ap "github.com/solidiquis/alacpretty_v2/internal"
	us "github.com/solidiquis/alacpretty_v2/internal/uistructs"
	ansi "github.com/solidiquis/ansigo"
	"os"
	"sort"
)

func themeShuffler(config []byte) {
	themes := make([]string, len(ap.Themes))

	i := 0
	for k, _ := range ap.Themes {
		themes[i] = ap.ThemePresenter(k)
		i++
	}

	sort.Strings(themes)

	curThm, err := ap.CurrentTheme(config)
	selected := 0

	if err == nil {
		curThm = ap.ThemePresenter(curThm)
		for i, thm := range themes {
			if thm == curThm {
				selected = i
				break
			}
		}
	}

	themesLi := us.InitList(themes, selected)
	themesLi.Render()

	ap.ListenForInput(config, themesLi)
}

func opacityGauge(config []byte, width, height, incrementer int) {
	curOp, err := ap.CurrentOpacity(config)
	if err != nil {
		panic(err)
	}

	g := us.InitGauge(int(curOp*100.0), width, height, incrementer)
	g.Render()

	stdin := make(chan string, 1)
	go ansi.GetChar(stdin)

	for {
		key := <-stdin
		g.Update(key)
		updatedConf := ap.ChangeOpacity(config, float64(g.Percent)/100.0)
		ap.ApplyChanges(updatedConf)
	}
}

func main() {
	config := ap.ConfigToBytes()

	switch os.Args[1] {
	case "theme_shuffler":
		themeShuffler(config)
	case "opacity_gauge":
		opacityGauge(config, 20, 1, 5)
	default:
		updatedConf := ap.ChangeOpacity(config, 1)
		ap.ApplyChanges(updatedConf)
	}
}
