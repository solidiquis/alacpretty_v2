package alacpretty

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
	"sort"
)

const (
	LIST_FRAME_LENGTH = 10
)

var (
	RIGHT_ARROW = ansi.FgBlue("\u25B6")
	DOWN_ARROW  = ansi.FgRed("\u25B2")
	UP_ARROW    = ansi.FgRed("\u25BC")
)

type List struct {
	FrameTail int
	Selected  int
	Items     []string
}

func InitList(items []string, selected int, frameLn int) *List {
	sort.Strings(items)

	return &List{
		FrameTail: frameLn,
		Selected:  selected,
		Items:     items,
	}
}

func (li *List) Render() {
	colOffset := "   "
	elements := make([]string, len(li.Items))

	for i, item := range li.Items {
		elements[i] = fmt.Sprintf("%s%s", colOffset, item)
	}

	elements[li.Selected] = fmt.Sprintf(
		" %s %s",
		RIGHT_ARROW,
		ansi.FgMagenta(li.Items[li.Selected]),
	)

	for i := li.FrameHead(); i < li.FrameTail; i++ {
		ansi.EraseLine()
		fmt.Println(elements[i])
	}
	ansi.CursorUp(LIST_FRAME_LENGTH)
}

func (li *List) Update(key string) {
	switch key {
	case "k", "<Up>":
		li.ShiftUp()
	case "j", "<Down>":
		li.ShiftDown()
	}
}

func (li *List) ShiftUp() {
	if li.Selected-1 < 0 {
		return
	}

	if li.Selected-1 < li.FrameHead() {
		li.FrameTail--
	}

	li.Selected--
	li.Render()
}

func (li *List) ShiftDown() {
	if li.Selected+1 > len(li.Items)-1 {
		return
	}

	if li.Selected+1 > li.FrameTail-1 {
		li.FrameTail++
	}

	li.Selected++
	li.Render()
}

func (li *List) FrameHead() int {
	return li.FrameTail - LIST_FRAME_LENGTH
}
