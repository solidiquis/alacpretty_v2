package alacpretty

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	RE_COLORS   = `\bcolors:.*(?:\n\s{2,}.+)+`
	RE_FONTSIZE = `size:\s*\d+\.0`
	RE_OPACITY  = `background_opacity:\s*\d+.\d+`
)

// Sets global variable ConfPath in init.go
func FindConf() {
	homeDir, err := os.UserHomeDir()
	must(err)

	// https://github.com/alacritty/alacritty
	paths := []string{
		".config/alacritty/alacritty.yml",
		".alacritty.yml",
		"alacritty/.alacritty.yml",
	}

	for _, path := range paths {
		tmp := fmt.Sprintf("%s/%s", homeDir, path)
		_, err := os.Stat(tmp)
		if err == nil {
			ConfPath = tmp
			return
		}
	}

	logErrExit(Errors["CONFIG_MISSING"])
}

// Returns path to alacritty configs.
func ConfigToBytes() []byte {
	conf, err := ioutil.ReadFile(ConfPath)
	must(err)

	return conf
}

func ApplyChanges(updatedConf string) {
	err := os.Truncate(ConfPath, 0)
	must(err)

	f, err := os.OpenFile(ConfPath, os.O_APPEND|os.O_WRONLY, 0644)
	must(err)
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(updatedConf)

	writer.Flush()
}

// Reports the current theme
func CurrentTheme(conf []byte) (string, error) {
	for k, v := range Themes {
		match, _ := regexp.Match(v, conf)
		if match {
			return k, nil
		}
	}

	return "", Errors["THEME_UNKNOWN"]
}

func ChangeTheme(config []byte, theme string) {
	newTh, ok := Themes[theme]
	if !ok {
		logErrExit(Errors["THEME_UNKNOWN"])
	}

	re, _ := regexp.Compile(RE_COLORS)
	if match := re.Match(config); !match {
		logErrExit(Errors["YAML_FORMAT"])
	}

	updatedConf := re.ReplaceAllString(string(config), newTh)

	ApplyChanges(updatedConf)
}
