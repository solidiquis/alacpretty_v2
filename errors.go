package alacpretty

import (
	"errors"
	ansi "github.com/solidiquis/ansigo"
	"log"
	"os"
)

var (
	ErrLog *log.Logger
	Errors = make(map[string]error)
)

func init() {
	ErrLog = log.New(os.Stderr, ansi.FgRed("ERROR:\t"), log.Ldate|log.Ltime|log.Lshortfile)

	Errors = map[string]error{
		"CONFIG_MISSING": errors.New("alacritty.yml not found."),
		"THEME_UNKNOWN":  errors.New("Unrecognized theme."),
		"THEME_INVALID":  errors.New("Argument 'theme' must be a valid key in Themes."),
		"YAML_FORMAT":    errors.New("Please make sure that the alacritty.yml is formatted properly."),
	}
}

func must(err error) {
	if err != nil {
		ErrLog.Fatalln(err)
	}
}

func logErrExit(err error) {
	ErrLog.Fatalln(err)
}
