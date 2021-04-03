package alacpretty

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
)

const (
	// Determines how many items to show before overflow.
	LIST_FRAME_LENGTH = 10
)

var (
	RIGHT_ARROW = ansi.FgBlue("\u25B6")
	DOWN_ARROW  = ansi.FgRed("\u25B2")
	UP_ARROW    = ansi.FgRed("\u25BC")
)

// List encapsulates everything required to render a scrollable list.
type List struct {
	// Tail end of the list's frame.
	FrameTail int

	// Index representing selected item in Items.
	Selected int

	// Items of the list.
	Items []string
}

// Initializes a List struct with items sorted alphabetically.
func InitList(items []string, selected int) *List {
	if selected > len(items)-1 {
		selected = len(items) - 1
	}

	frameLn := LIST_FRAME_LENGTH
	offset := LIST_FRAME_LENGTH - 4

	if selected > frameLn-1 {
		if selected+offset > len(items)-1 {
			frameLn = len(items)
		} else {
			frameLn = selected + offset
		}
	}

	return &List{
		FrameTail: frameLn,
		Selected:  selected,
		Items:     items,
	}
}

// Renders list to the terminal, LIST_FRAME_LENGTH items at a time.
// Uses a sliding window determined by li.FrameTail and li.FrameHead()
// to determine which items are visible at any given point in time.
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

// Implements scrolling behavior of list, allowing user to navigate
// through the list of items with the "k", "j", up, and down keys.
func (li *List) Update(key string) {
	switch key {
	case "k", "<Up>":
		li.ShiftUp()
	case "j", "<Down>":
		li.ShiftDown()
	}
}

// Handles upwards list navigation.
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

// Handles downwards list navigation.
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

// Determines the starting, leftmost position of the sliding window.
func (li *List) FrameHead() int {
	return li.FrameTail - LIST_FRAME_LENGTH
}
