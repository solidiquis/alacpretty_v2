package alacpretty

var ConfPath string

const (
	KEY_RETURN = 0xa
	KEY_ESC    = 0x1b
)

func init() {
	FindConf()
}
