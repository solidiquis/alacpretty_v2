package alacpretty

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
	"math"
	"strings"
)

const (
	TL_VRTX = "\u250C"
	TR_VRTX = "\u2510"
	BL_VRTX = "\u2514"
	BR_VRTX = "\u2518"
	HZ_LINE = "\u2500"
	VT_LINE = "\u2502"
	BX_SHD  = "\u258A"
)

// Gauge encapsulates everything required to create an adjustable gauge.
type Gauge struct {
	// How tall and wide is the gauge.
	Height, Width int

	// Current percent filled.
	Percent int

	// How much to inc/decrement gauge by.
	Incrementer int
}

func InitGauge(percent, width, height, incrementer int) *Gauge {
	if percent > 100 {
		percent = 100
	}

	return &Gauge{
		Percent:     percent,
		Width:       width,
		Height:      height,
		Incrementer: incrementer,
	}
}

func (g *Gauge) Render() {
	mdSps := make([]string, g.Width)
	hzLns := make([]string, g.Width)

	for i := 0; i < g.Width; i++ {
		hzLns[i] = HZ_LINE

		if float64(i+1)/float64(g.Width)*float64(100) <= float64(g.Percent) {
			mdSps[i] = ansi.FgMagenta(ansi.BgMagenta(BX_SHD))
		} else {
			mdSps[i] = " "
		}
	}

	midpt := g.Width / 2
	percentStr := g.PercentStr()
	startPos := midpt - int(math.Round(float64(len(percentStr))/2.0))

	percentTxt := mdSps
	for _, ch := range percentStr {
		if startPos*5 <= g.Percent {
			percentTxt[startPos] = ansi.BgMagenta(ansi.FgBlack(string(ch)))
		} else {
			percentTxt[startPos] = string(ch)
		}
		startPos++
	}

	hzLn := strings.Join(hzLns, "")
	mdSp := strings.Join(mdSps, "")

	midSec := make([]string, g.Height)
	for i := 0; i < g.Height; i++ {
		midSec[i] = fmt.Sprint(VT_LINE, mdSp, VT_LINE)
	}

	htMidpt := int(math.Round(float64(g.Height)/2.0)) - 1
	midSec[htMidpt] = fmt.Sprint(VT_LINE, strings.Join(percentTxt, ""), VT_LINE)

	top := fmt.Sprint(TL_VRTX, hzLn, TR_VRTX, "\n")
	btm := fmt.Sprint(BL_VRTX, hzLn, BR_VRTX, "\n")

	fmt.Print(top)
	for _, mid := range midSec {
		fmt.Println(mid)
	}
	fmt.Print(btm)
}

func (g *Gauge) Update(key string) {
	switch key {
	case "h", "<Left>":
		g.Decrement()
	case "l", "<Right>":
		g.Increment()
	}
}

func (g *Gauge) PercentStr() string {
	return fmt.Sprintf("%.1f%%", float64(g.Percent))
}

func (g *Gauge) Increment() {
	if g.Percent >= 100 {
		return
	}

	if g.Percent+g.Incrementer > 100 {
		g.Percent = 100
	} else {
		g.Percent += g.Incrementer
	}

	ansi.CursorUp(g.Height + 2)
	g.Render()
}

func (g *Gauge) Decrement() {
	if g.Percent <= 0 {
		return
	}

	if g.Percent-g.Incrementer < 0 {
		g.Percent = 0
	} else {
		g.Percent -= 5
	}

	ansi.CursorUp(g.Height + 2)
	g.Render()
}
