package alacpretty

import (
	"fmt"
	"log"
	"os"
)

// Write logs to a separate terminal
func Debug(charDevice string) func([]byte) {
	dev := fmt.Sprintf("/dev/%s", charDevice)
	f, err := os.OpenFile(dev, os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	f.Write([]byte("\n"))

	return func(b []byte) {
		f.Write([]byte(b))
	}
}
