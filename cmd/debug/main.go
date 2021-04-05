// For manual testing of ui components in isolation.
package main

import (
	"fmt"
	"github.com/sahilm/fuzzy"
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

	//curThm, err := ap.CurrentTheme(config)
	selected := 0

	//if err == nil {
	//curThm = ap.ThemePresenter(curThm)
	//for i, thm := range themes {
	//if thm == curThm {
	//selected = i
	//break
	//}
	//}
	//}

	for i := 0; i < 11; i++ {
		fmt.Println()
	}
	ansi.CursorUp(11)

	li := us.InitList(themes, selected)

	stdin := make(chan string, 1)
	go ansi.GetChar(stdin)
	input := ""

	fmt.Printf("%s: ", ansi.Bright("Theme"))
	ansi.CursorSavePos()
	fmt.Println()
	li.Render()
	ansi.CursorRestorePos()

	for {
		switch key := <-stdin; key {
		case "<Backspace>":
			if len(input) > 0 {
				ansi.Backspace()
				input = input[:len(input)-1]
			}
		case "<Down>":
			fmt.Println()
			li.ShiftDown()
			ansi.CursorRestorePos()
			continue
		case "<Up>":
			fmt.Println()
			li.ShiftUp()
			ansi.CursorRestorePos()
			continue
		case "<Enter>":
			theme := li.MakeSelection()
			theme = ap.FormatTheme(theme)
			updatedConf := ap.ChangeTheme(config, theme)
			ap.ApplyChanges(updatedConf)
		default:
			li.Selected = 0
			fmt.Print(key)
			input += key
		}

		ansi.CursorSavePos()
		fmt.Println()

		matches := fuzzy.Find(input, li.Items)
		items := make([]string, len(li.Items))
		for i, match := range matches {
			items[i] = match.Str
		}

		if items[0] != "" {
			li.Matches = items
			li.FrameTail = 10 // make better
			li.Render()
		} else {
			li.Matches = nil
			li.Render()
		}
		ansi.CursorRestorePos()
	}
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
