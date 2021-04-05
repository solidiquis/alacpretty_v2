package alacpretty

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
)

const (
	// Determines how many items to show before overflow.
	OVERFLOW_LENGTH = 10
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

	// Items that match user input.
	Matches []string

	// User input.
	Input string
}

// Initializes a List struct with items sorted alphabetically.
func InitList(items []string, selected int) *List {
	if selected > len(items)-1 {
		selected = len(items) - 1
	}

	frameLn := len(items)
	offset := 0

	if len(items) > OVERFLOW_LENGTH {
		frameLn = OVERFLOW_LENGTH
		offset = OVERFLOW_LENGTH - 4
	}

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

// Renders list to the terminal, OVERFLOW_LENGTH items at a time.
// Uses a sliding window determined by li.FrameTail and li.FrameHead()
// to determine which items are visible at any given point in time.
func (li *List) Render() {
	colOffset := "   "

	var elements []string
	var iterable []string

	if len(li.Matches) > 0 {
		elements = make([]string, len(li.Matches))
		iterable = li.Matches
	} else {
		elements = make([]string, len(li.Items))
		iterable = li.Items
	}

	for i, item := range iterable {
		elements[i] = fmt.Sprintf("%s%s", colOffset, item)
	}

	elements[li.Selected] = fmt.Sprintf(
		" %s %s",
		RIGHT_ARROW,
		ansi.FgMagenta(iterable[li.Selected]),
	)

	for i := li.FrameHead(); i < li.FrameTail; i++ {
		ansi.EraseLine()
		fmt.Println(elements[i])
	}
	ansi.CursorUp(len(iterable))
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
	if li.Selected+1 > len(li.Items)-1 || li.NextItem() == "" {
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
	return li.FrameTail - OVERFLOW_LENGTH
}

func (li *List) MakeSelection() string {
	if len(li.Matches) > 0 {
		return li.Matches[li.Selected]
	}

	return li.Items[li.Selected]
}

func (li *List) FocusedItem() string {
	if len(li.Matches) > 0 {
		return li.Matches[li.Selected]
	}

	return li.Items[li.Selected]
}

func (li *List) NextItem() string {
	if len(li.Matches) > 0 {
		return li.Matches[li.Selected+1]
	}

	return li.Items[li.Selected+1]
}
