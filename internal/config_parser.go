package alacpretty

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	RE_COLORS   = `\bcolors:.*(?:\n\s{2,}.+)+`
	RE_FONTSIZE = `size:\s*\d+\.0`
	RE_OPACITY  = `background_opacity:\s*\d+.\d+`
)

// Sets global variable ConfPath in init.go.
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

// Truncates alacritty.yml and replaces it with updatedConf
func ApplyChanges(updatedConf []byte) {
	err := os.Truncate(ConfPath, 0)
	must(err)

	f, err := os.OpenFile(ConfPath, os.O_APPEND|os.O_WRONLY, 0644)
	must(err)
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(string((updatedConf)))

	writer.Flush()
}

// Reports the current theme.
func CurrentTheme(conf []byte) (string, error) {
	for k, v := range Themes {
		match, _ := regexp.Match(v, conf)
		if match {
			return k, nil
		}
	}

	return "", Errors["THEME_UNKNOWN"]
}

// Returns config as byte slice with updated theme.
func ChangeTheme(config []byte, theme string) []byte {
	newTh, ok := Themes[theme]
	if !ok {
		logErrExit(Errors["THEME_UNKNOWN"])
	}

	newTh = strings.Trim(newTh, "\n")

	re, _ := regexp.Compile(RE_COLORS)
	if match := re.Match(config); !match {
		logErrExit(Errors["YAML_FORMAT"])
	}

	updatedConf := re.ReplaceAllString(string(config), newTh)

	return []byte(updatedConf)
}

// Returns the current opacity.
//func CurrentOpacity() float64 {

//}
