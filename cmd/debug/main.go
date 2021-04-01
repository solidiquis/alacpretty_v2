package main

import (
	ap "github.com/solidiquis/alacpretty_v2/internal"
)

func main() {
	config := ap.ConfigToBytes()
	ap.ChangeTheme(config, "SOLARIZED_DARK")
}
