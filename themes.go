package alacpretty

import (
	"strings"
)

func ThemePresenter(txt string) string {
	return strings.Title(strings.ToLower(strings.Replace(txt, "_", " ", -1)))
}

func FormatTheme(txt string) string {
	return strings.ToUpper(strings.Replace(txt, " ", "_", -1))
}

var Themes = map[string]string{
	"AFTERGLOW":             AFTERGLOW,
	"ARGONAUT":              ARGONAUT,
	"AYU_DARK":              AYU_DARK,
	"BASE16_DEFAULT_DARK":   BASE16_DEFAULT_DARK,
	"BLOOD_MOON":            BLOOD_MOON,
	"BREEZE":                BREEZE,
	"CAMPBELL":              CAMPBELL,
	"CHALLENGER_DEEP":       CHALLENGER_DEEP,
	"COBALT2":               COBALT2,
	"CYBERPUNK_NEON":        CYBERPUNK_NEON,
	"DARCULA":               DARCULA,
	"DARK_PASTELS":          DARK_PASTELS,
	"DRACULA":               DRACULA,
	"FALCON":                FALCON,
	"GOTHAM":                GOTHAM,
	"GRUVBOX_DARK":          GRUVBOX_DARK,
	"GRUVBOX_LIGHT":         GRUVBOX_LIGHT,
	"HIGH_CONTRAST":         HIGH_CONTRAST,
	"HYPER":                 HYPER,
	"ITERM":                 ITERM,
	"KONSOLE_LINUX":         KONSOLE_LINUX,
	"LOW_CONTRAST":          LOW_CONTRAST,
	"MATERIAL_THEME":        MATERIAL_THEME,
	"MATERIAL_THEME_MOD":    MATERIAL_THEME_MOD,
	"NORD":                  NORD,
	"OCEANIC_NEXT":          OCEANIC_NEXT,
	"ONE_DARK":              ONE_DARK,
	"PAPERCOLOR_LIGHT":      PAPERCOLOR_LIGHT,
	"PENCIL_DARK":           PENCIL_DARK,
	"PENCIL_LIGHT":          PENCIL_LIGHT,
	"REMEDY_DARK":           REMEDY_DARK,
	"SNAZZY":                SNAZZY,
	"SOLARIZED_DARK":        SOLARIZED_DARK,
	"SOLARIZED_LIGHT":       SOLARIZED_LIGHT,
	"STORM":                 STORM,
	"TAERMINAL":             TAERMINAL,
	"TANGO_DARK":            TANGO_DARK,
	"TENDER":                TENDER,
	"TERMINAL_APP":          TERMINAL_APP,
	"THE_LOVELACE":          THE_LOVELACE,
	"TOMORROW_NIGHT":        TOMORROW_NIGHT,
	"TOMORROW_NIGHT_BRIGHT": TOMORROW_NIGHT_BRIGHT,
	"WOMBAT":                WOMBAT,
	"XTERM":                 XTERM,
}

const COBALT2 = `
# From the famous Cobalt2 sublime theme
# Source : https://github.com/wesbos/cobalt2/tree/master/Cobalt2
colors:
  # Default colors
  primary:
    background: '0x122637'
    foreground: '0xffffff'

  cursor:
    text: '0x122637'
    cursor: '0xf0cb09'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xff0000'
    green:   '0x37dd21'
    yellow:  '0xfee409'
    blue:    '0x1460d2'
    magenta: '0xff005d'
    cyan:    '0x00bbbb'
    white:   '0xbbbbbb'

  # Bright colors
  bright:
    black:   '0x545454'
    red:     '0xf40d17'
    green:   '0x3bcf1d'
    yellow:  '0xecc809'
    blue:    '0x5555ff'
    magenta: '0xff55ff'
    cyan:    '0x6ae3f9'
    white:   '0xffffff'
`

const AFTERGLOW = `
colors:
  # Default colors
  primary:
    background: '0x2c2c2c'
    foreground: '0xd6d6d6'

    dim_foreground:    '0xdbdbdb'
    bright_foreground: '0xd9d9d9'
    dim_background:    '0x202020' # not sure
    bright_background: '0x3a3a3a' # not sure

  # Cursor colors
  cursor:
    text:   '0x2c2c2c'
    cursor: '0xd9d9d9'

  # Normal colors
  normal:
    black:   '0x1c1c1c'
    red:     '0xbc5653'
    green:   '0x909d63'
    yellow:  '0xebc17a'
    blue:    '0x7eaac7'
    magenta: '0xaa6292'
    cyan:    '0x86d3ce'
    white:   '0xcacaca'

  # Bright colors
  bright:
    black:   '0x636363'
    red:     '0xbc5653'
    green:   '0x909d63'
    yellow:  '0xebc17a'
    blue:    '0x7eaac7'
    magenta: '0xaa6292'
    cyan:    '0x86d3ce'
    white:   '0xf7f7f7'

  # Dim colors
  dim:
    black:   '0x232323'
    red:     '0x74423f'
    green:   '0x5e6547'
    yellow:  '0x8b7653'
    blue:    '0x556b79'
    magenta: '0x6e4962'
    cyan:    '0x5c8482'
    white:   '0x828282'
`

const ARGONAUT = `
colors:
  # Default colors
  primary:
    background: '0x292C3E'
    foreground: '0xEBEBEB'

  # Cursor colors
  cursor:
    text: '0xFF261E'
    cursor: '0xFF261E'

  # Normal colors
  normal:
    black:   '0x0d0d0d'
    red:     '0xFF301B'
    green:   '0xA0E521'
    yellow:  '0xFFC620'
    blue:    '0x1BA6FA'
    magenta: '0x8763B8'
    cyan:    '0x21DEEF'
    white:   '0xEBEBEB'

  # Bright colors
  bright:
    black:   '0x6D7070'
    red:     '0xFF4352'
    green:   '0xB8E466'
    yellow:  '0xFFD750'
    blue:    '0x1BA6FA'
    magenta: '0xA578EA'
    cyan:    '0x73FBF1'
    white:   '0xFEFEF8'
`

const AYU_DARK = `
# Colors (Ayu Dark)
colors:
  # Default colors
  primary:
    background: '0x0A0E14'
    foreground: '0xB3B1AD'

  # Normal colors
  normal:
    black:   '0x01060E'
    red:     '0xEA6C73'
    green:   '0x91B362'
    yellow:  '0xF9AF4F'
    blue:    '0x53BDFA'
    magenta: '0xFAE994'
    cyan:    '0x90E1C6'
    white:   '0xC7C7C7'

  # Bright colors
  bright:
    black:   '0x686868'
    red:     '0xF07178'
    green:   '0xC2D94C'
    yellow:  '0xFFB454'
    blue:    '0x59C2FF'
    magenta: '0xFFEE99'
    cyan:    '0x95E6CB'
    white:   '0xFFFFFF'
`

const BASE16_DEFAULT_DARK = `
# Colors (Base16 Default Dark)
colors:
  # Default colors
  primary:
    background: '0x181818'
    foreground: '0xd8d8d8'

  cursor:
    text: '0xd8d8d8'
    cursor: '0xd8d8d8'

  # Normal colors
  normal:
    black:   '0x181818'
    red:     '0xab4642'
    green:   '0xa1b56c'
    yellow:  '0xf7ca88'
    blue:    '0x7cafc2'
    magenta: '0xba8baf'
    cyan:    '0x86c1b9'
    white:   '0xd8d8d8'

  # Bright colors
  bright:
    black:   '0x585858'
    red:     '0xab4642'
    green:   '0xa1b56c'
    yellow:  '0xf7ca88'
    blue:    '0x7cafc2'
    magenta: '0xba8baf'
    cyan:    '0x86c1b9'
    white:   '0xf8f8f8'
`

const BLOOD_MOON = `
# Colors (Blood Moon)
colors:
  # Default colors
  primary:
    background: '0x10100E'
    foreground: '0xC6C6C4'

  # Normal colors
  normal:
    black:   '0x10100E'
    red:     '0xC40233'
    green:   '0x009F6B'
    yellow:  '0xFFD700'
    blue:    '0x0087BD'
    magenta: '0x9A4EAE'
    cyan:    '0x20B2AA'
    white:   '0xC6C6C4'

  # Bright colors
  bright:
    black:   '0x696969'
    red:     '0xFF2400'
    green:   '0x03C03C'
    yellow:  '0xFDFF00'
    blue:    '0x007FFF'
    magenta: '0xFF1493'
    cyan:    '0x00CCCC'
    white:   '0xFFFAFA'
`

const BREEZE = `
# KDE Breeze (Ported from Konsole)
colors:
  # Default colors
  primary:
    background: '0x232627'
    foreground: '0xfcfcfc'

    dim_foreground: '0xeff0f1'
    bright_foreground: '0xffffff'
    dim_background: '0x31363b'
    bright_background: '0x000000'

  # Normal colors
  normal:
    black: '0x232627'
    red: '0xed1515'
    green: '0x11d116'
    yellow: '0xf67400'
    blue: '0x1d99f3'
    magenta: '0x9b59b6'
    cyan: '0x1abc9c'
    white: '0xfcfcfc'

  # Bright colors
  bright:
    black: '0x7f8c8d'
    red: '0xc0392b'
    green: '0x1cdc9a'
    yellow: '0xfdbc4b'
    blue: '0x3daee9'
    magenta: '0x8e44ad'
    cyan: '0x16a085'
    white: '0xffffff'

  # Dim colors
  dim:
    black: '0x31363b'
    red: '0x783228'
    green: '0x17a262'
    yellow: '0xb65619'
    blue: '0x1b668f'
    magenta: '0x614a73'
    cyan: '0x186c60'
    white: '0x63686d'
`

const CAMPBELL = `
# Campbell (Windows 10 default)
colors:
  # Default colors
  primary:
    background: '0x0c0c0c'
    foreground: '0xcccccc'

  # Normal colors
  normal:
    black:      '0x0c0c0c'
    red:        '0xc50f1f'
    green:      '0x13a10e'
    yellow:     '0xc19c00'
    blue:       '0x0037da'
    magenta:    '0x881798'
    cyan:       '0x3a96dd'
    white:      '0xcccccc'

  # Bright colors
  bright:
    black:      '0x767676'
    red:        '0xe74856'
    green:      '0x16c60c'
    yellow:     '0xf9f1a5'
    blue:       '0x3b78ff'
    magenta:    '0xb4009e'
    cyan:       '0x61d6d6'
    white:      '0xf2f2f2'
`

const CHALLENGER_DEEP = `
# Colors (Challenger Deep)
colors:
  # Default colors
  primary:
    background: '0x1e1c31'
    foreground: '0xcbe1e7'

  cursor:
    text: '0xff271d'
    cursor: '0xfbfcfc'
  # Normal colors
  normal:
    black:   '0x141228'
    red:     '0xff5458'
    green:   '0x62d196'
    yellow:  '0xffb378'
    blue:    '0x65b2ff'
    magenta: '0x906cff'
    cyan:    '0x63f2f1'
    white:   '0xa6b3cc'
  # Bright colors
  bright:
    black:   '0x565575'
    red:     '0xff8080'
    green:   '0x95ffa4'
    yellow:  '0xffe9aa'
    blue:    '0x91ddff'
    magenta: '0xc991e1'
    cyan:    '0xaaffe4'
    white:   '0xcbe3e7'
`

const CYBERPUNK_NEON = `
# Cyber Punk Neon
# Source: https://github.com/Roboron3042/Cyberpunk-Neon
colors:
  # Default colors
  primary:
    background: "0x000b1e"
    foreground: "0x0abdc6"

  cursor:
    text:   "0x000b1e"
    cursor: "0x0abdc6"

  # Normal colors
  normal:
    black:   "0x123e7c"
    red:     "0xff0000"
    green:   "0xd300c4"
    yellow:  "0xf57800"
    blue:    "0x123e7c"
    magenta: "0x711c91"
    cyan:    "0x0abdc6"
    white:   "0xd7d7d5"

  # Bright colors
  bright:
    black:   "0x1c61c2"
    red:     "0xff0000"
    green:   "0xd300c4"
    yellow:  "0xf57800"
    blue:    "0x00ff00"
    magenta: "0x711c91"
    cyan:    "0x0abdc6"
    white:   "0xd7d7d5"
`

const DARCULA = `
# Colors (Dracula)
colors:
  # Default colors
  primary:
    background: '0x282a36'
    foreground: '0xf8f8f2'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xff5555'
    green:   '0x50fa7b'
    yellow:  '0xf1fa8c'
    blue:    '0xcaa9fa'
    magenta: '0xff79c6'
    cyan:    '0x8be9fd'
    white:   '0xbfbfbf'

  # Bright colors
  bright:
    black:   '0x282a35'
    red:     '0xff6e67'
    green:   '0x5af78e'
    yellow:  '0xf4f99d'
    blue:    '0xcaa9fa'
    magenta: '0xff92d0'
    cyan:    '0x9aedfe'
    white:   '0xe6e6e6'
`

const DARK_PASTELS = `
# Colors (Konsole's Dark Pastels)
colors:
  # Default colors
  primary:
    background: '0x2C2C2C'
    foreground: '0xDCDCCC'

  # Normal colors
  normal:
    black:   '0x3F3F3F'
    red:     '0x705050'
    green:   '0x60B48A'
    yellow:  '0xDFAF8F'
    blue:    '0x9AB8D7'
    magenta: '0xDC8CC3'
    cyan:    '0x8CD0D3'
    white:   '0xDCDCCC'

  # Bright colors
  bright:
    black:   '0x709080'
    red:     '0xDCA3A3'
    green:   '0x72D5A3'
    yellow:  '0xF0DFAF'
    blue:    '0x94BFF3'
    magenta: '0xEC93D3'
    cyan:    '0x93E0E3'
    white:   '0xFFFFFF'
`

const DRACULA = `
# Colors (Dracula)
colors:
  # Default colors
  primary:
    background: '0x282a36'
    foreground: '0xf8f8f2'
 
  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xff5555'
    green:   '0x50fa7b'
    yellow:  '0xf1fa8c'
    blue:    '0xbd93f9'
    magenta: '0xff79c6'
    cyan:    '0x8be9fd'
    white:   '0xbbbbbb'
 
  # Bright colors
  bright:
    black:   '0x555555'
    red:     '0xff5555'
    green:   '0x50fa7b'
    yellow:  '0xf1fa8c'
    blue:    '0xcaa9fa'
    magenta: '0xff79c6'
    cyan:    '0x8be9fd'
    white:   '0xffffff'
`

const FALCON = `
# falcon colorscheme for alacritty
# by fenetikm, https://github.com/fenetikm/falcon
colors:
  # Default colors
  primary:
    background: '0x020221'
    foreground: '0xb4b4b9'

  cursor:
    text: '0x020221'
    cursor: '0xffe8c0'

  # Normal colors
  normal:
    black:   '0x000004'
    red:     '0xff3600'
    green:   '0x718e3f'
    yellow:  '0xffc552'
    blue:    '0x635196'
    magenta: '0xff761a'
    cyan:    '0x34bfa4'
    white:   '0xb4b4b9'

  # Bright colors
  bright:
    black:   '0x020221'
    red:     '0xff8e78'
    green:   '0xb1bf75'
    yellow:  '0xffd392'
    blue:    '0x99a4bc'
    magenta: '0xffb07b'
    cyan:    '0x8bccbf'
    white:   '0xf8f8ff'
`

const GOTHAM = `
# Colors (Gotham)
colors:
  # Default colors
  primary:
    background: '0x0a0f14'
    foreground: '0x98d1ce'

  # Normal colors
  normal:
    black: '0x0a0f14'
    red: '0xc33027'
    green: '0x26a98b'
    yellow: '0xedb54b'
    blue: '0x195465'
    magenta: '0x4e5165'
    cyan: '0x33859d'
    white: '0x98d1ce'

  # Bright colors
  bright:
    black: '0x10151b'
    red: '0xd26939'
    green: '0x081f2d'
    yellow: '0x245361'
    blue: '0x093748'
    magenta: '0x888ba5'
    cyan: '0x599caa'
    white: '0xd3ebe9'
`

const GRUVBOX_DARK = `
# Colors (Gruvbox dark)
colors:
  # Default colors
  primary:
    # hard contrast: background = '0x1d2021'
    background: '0x282828'
    # soft contrast: background = '0x32302f'
    foreground: '0xebdbb2'

  # Normal colors
  normal:
    black:   '0x282828'
    red:     '0xcc241d'
    green:   '0x98971a'
    yellow:  '0xd79921'
    blue:    '0x458588'
    magenta: '0xb16286'
    cyan:    '0x689d6a'
    white:   '0xa89984'

  # Bright colors
  bright:
    black:   '0x928374'
    red:     '0xfb4934'
    green:   '0xb8bb26'
    yellow:  '0xfabd2f'
    blue:    '0x83a598'
    magenta: '0xd3869b'
    cyan:    '0x8ec07c'
    white:   '0xebdbb2'
`

const GRUVBOX_LIGHT = `
# Colors (Gruvbox light)
colors:
  # Default colors
  primary:
    # hard contrast: background = '0xf9f5d7'
    background: '0xfbf1c7'
    # soft contrast: background = '0xf2e5bc'
    foreground: '0x3c3836'

  # Normal colors
  normal:
    black:   '0xfbf1c7'
    red:     '0xcc241d'
    green:   '0x98971a'
    yellow:  '0xd79921'
    blue:    '0x458588'
    magenta: '0xb16286'
    cyan:    '0x689d6a'
    white:   '0x7c6f64'

  # Bright colors
  bright:
    black:   '0x928374'
    red:     '0x9d0006'
    green:   '0x79740e'
    yellow:  '0xb57614'
    blue:    '0x076678'
    magenta: '0x8f3f71'
    cyan:    '0x427b58'
    white:   '0x3c3836'
`

const HIGH_CONTRAST = `
# Colors (High Contrast)
colors:
  # Default colors
  primary:
    background: '0x444444'
    foreground: '0xdddddd'

  cursor:
    text: '0xaaaaaa'
    cursor: '0xffffff'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xff0000'
    green:   '0x00ff00'
    yellow:  '0xffff00'
    blue:    '0x0000ff'
    magenta: '0xff00ff'
    cyan:    '0x00ffff'
    white:   '0xffffff'

  # Bright colors
  bright:
    black:   '0x000000'
    red:     '0xff0000'
    green:   '0x00ff00'
    yellow:  '0xffff00'
    blue:    '0x0000ff'
    magenta: '0xff00ff'
    cyan:    '0x00ffff'
    white:   '0xffffff'
`

const HYPER = `
# Colors (Hyper)
colors:
  # Default colors
  primary:
    background: '0x000000'
    foreground: '0xffffff'
  cursor:
    text: '0xF81CE5'
    cursor: '0xffffff'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xfe0100'
    green:   '0x33ff00'
    yellow:  '0xfeff00'
    blue:    '0x0066ff'
    magenta: '0xcc00ff'
    cyan:    '0x00ffff'
    white:   '0xd0d0d0'

  # Bright colors
  bright:
    black:   '0x808080'
    red:     '0xfe0100'
    green:   '0x33ff00'
    yellow:  '0xfeff00'
    blue:    '0x0066ff'
    magenta: '0xcc00ff'
    cyan:    '0x00ffff'
    white:   '0xFFFFFF'
`

const ITERM = `
# Colors (iTerm 2 default theme)
colors:
  # Default colors
  primary:
    background: '0x101421'
    foreground: '0xfffbf6'

  # Normal colors
  normal:
    black:   '0x2e2e2e'
    red:     '0xeb4129'
    green:   '0xabe047'
    yellow:  '0xf6c744'
    blue:    '0x47a0f3'
    magenta: '0x7b5cb0'
    cyan:    '0x64dbed'
    white:   '0xe5e9f0'

  # Bright colors
  bright:
    black:   '0x565656'
    red:     '0xec5357'
    green:   '0xc0e17d'
    yellow:  '0xf9da6a'
    blue:    '0x49a4f8'
    magenta: '0xa47de9'
    cyan:    '0x99faf2'
    white:   '0xffffff'
`

const KONSOLE_LINUX = `
# Color theme ported from Konsole: Linux colors 
colors:  
  primary:
    foreground: '0xe3e3e3'
    bright_foreground: '0xffffff'
    dim_foreground:    '0xe3e3e3'
    background: '0x1f1f1f'
    bright_background: '0x686868' # not sure
    dim_background:    '0x1f1f1f' # not sure

  cursor:
    text: '0x191622'
    cursor: '0xf8f8f2'

  search:
    matches:
      foreground: '0xb2b2b2'
      background: '0xb26818'
    focused_match:
      foreground: CellBackground
      background: CellForeground

  normal:
    black:   '0x000000'
    red:     '0xb21818'
    green:   '0x18b218'
    yellow:  '0xb26818'
    blue:    '0x1818b2'
    magenta: '0xb218b2'
    cyan:    '0x18b2b2'
    white:   '0xb2b2b2'

  bright:
    black:   '0x686868'
    red:     '0xff5454'
    green:   '0x54ff54'
    yellow:  '0xffff54'
    blue:    '0x5454ff'
    magenta: '0xff54ff'
    cyan:    '0x54ffff'
    white:   '0xffffff'

  dim:
    black:   '0x000000'
    red:     '0xb21818'
    green:   '0x18b218'
    yellow:  '0xb26818'
    blue:    '0x1818b2'
    magenta: '0xb218b2'
    cyan:    '0x18b2b2'
    white:   '0xb2b2b2'
`

const LOW_CONTRAST = `
# Colors (Dim)
colors:
  # Default colors
  primary:
    background: '0x333333'
    foreground: '0xdddddd'

  cursor:
    text: '0xaaaaaa'
    cursor: '0xffffff'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xbb0000'
    green:   '0x00bb00'
    yellow:  '0xbbbb00'
    blue:    '0x0000bb'
    magenta: '0xbb00bb'
    cyan:    '0x00bbbb'
    white:   '0xbbbbbb'

  # Bright colors
  bright:
    black:   '0x000000'
    red:     '0xbb0000'
    green:   '0x00bb00'
    yellow:  '0xbbbb00'
    blue:    '0x0000bb'
    magenta: '0xbb00bb'
    cyan:    '0x00bbbb'
    white:   '0xbbbbbb'
`

const MATERIAL_THEME = `
# Colors (Material Theme)
colors:
  # Default colors
  primary:
    background: '0x1e282d'
    foreground: '0xc4c7d1'

  # Normal colors
  normal:
    black:   '0x666666'
    red:     '0xeb606b'
    green:   '0xc3e88d'
    yellow:  '0xf7eb95'
    blue:    '0x80cbc4'
    magenta: '0xff2f90'
    cyan:    '0xaeddff'
    white:   '0xffffff'

  # Bright colors
  bright:
    black:   '0xff262b'
    red:     '0xeb606b'
    green:   '0xc3e88d'
    yellow:  '0xf7eb95'
    blue:    '0x7dc6bf'
    magenta: '0x6c71c4'
    cyan:    '0x35434d'
    white:   '0xffffff'
`

const MATERIAL_THEME_MOD = `
# Colors (Material Theme)
colors:
  # Default colors
  primary:
    background: '0x1e282d'
    foreground: '0xc4c7d1'

  # Normal colors
  normal:
    black:   '0x666666'
    red:     '0xeb606b'
    green:   '0xc3e88d'
    yellow:  '0xf7eb95'
    blue:    '0x80cbc4'
    magenta: '0xff2f90'
    cyan:    '0xaeddff'
    white:   '0xffffff'

  # Bright colors
  bright:
    black:   '0xa1a1a1'
    red:     '0xeb606b'
    green:   '0xc3e88d'
    yellow:  '0xf7eb95'
    blue:    '0x7dc6bf'
    magenta: '0x6c71c4'
    cyan:    '0x35434d'
    white:   '0xffffff'
`

const NORD = `
# Colors (Nord)
colors:
  # Default colors
  primary:
    background: '0x2E3440'
    foreground: '0xD8DEE9'

  # Normal colors
  normal:
    black:   '0x3B4252'
    red:     '0xBF616A'
    green:   '0xA3BE8C'
    yellow:  '0xEBCB8B'
    blue:    '0x81A1C1'
    magenta: '0xB48EAD'
    cyan:    '0x88C0D0'
    white:   '0xE5E9F0'

  # Bright colors
  bright:
    black:   '0x4C566A'
    red:     '0xBF616A'
    green:   '0xA3BE8C'
    yellow:  '0xEBCB8B'
    blue:    '0x81A1C1'
    magenta: '0xB48EAD'
    cyan:    '0x8FBCBB'
    white:   '0xECEFF4'
`

const OCEANIC_NEXT = `
# Colors (Oceanic Next)
colors:
  # Default colors
  primary:
    background: '0x1b2b34'
    foreground: '0xd8dee9'

  # Normal colors
  normal:
    black:   '0x29414f'
    red:     '0xec5f67'
    green:   '0x99c794'
    yellow:  '0xfac863'
    blue:    '0x6699cc'
    magenta: '0xc594c5'
    cyan:    '0x5fb3b3'
    white:   '0x65737e'

  # Bright colors
  bright:
    black:   '0x405860'
    red:     '0xec5f67'
    green:   '0x99c794'
    yellow:  '0xfac863'
    blue:    '0x6699cc'
    magenta: '0xc594c5'
    cyan:    '0x5fb3b3'
    white:   '0xadb5c0'
`

const ONE_DARK = `
# Colors (One Dark)
colors:
  # Default colors
  primary:
    background: '0x1e2127'
    foreground: '0xabb2bf'

  # Normal colors
  normal:
    black:   '0x1e2127'
    red:     '0xe06c75'
    green:   '0x98c379'
    yellow:  '0xd19a66'
    blue:    '0x61afef'
    magenta: '0xc678dd'
    cyan:    '0x56b6c2'
    white:   '0xabb2bf'

  # Bright colors
  bright:
    black:   '0x5c6370'
    red:     '0xe06c75'
    green:   '0x98c379'
    yellow:  '0xd19a66'
    blue:    '0x61afef'
    magenta: '0xc678dd'
    cyan:    '0x56b6c2'
    white:   '0xffffff'
`

const PAPERCOLOR_LIGHT = `
# Colors (PaperColor - Light)
colors:
  # Default colors
  primary:
    background: '0xeeeeee'
    foreground: '0x878787'

  cursor:
    text: '0xeeeeee'
    cursor: '0x878787'

  # Normal colors
  normal:
    black:   '0xeeeeee'
    red:     '0xaf0000'
    green:   '0x008700'
    yellow:  '0x5f8700'
    blue:    '0x0087af'
    magenta: '0x878787'
    cyan:    '0x005f87'
    white:   '0x444444'

  # Bright colors
  bright:
    black:   '0xbcbcbc'
    red:     '0xd70000'
    green:   '0xd70087'
    yellow:  '0x8700af'
    blue:    '0xd75f00'
    magenta: '0xd75f00'
    cyan:    '0x005faf'
    white:   '0x005f87'
`

const PENCIL_DARK = `
# Colors (Pencil Dark)
colors:
  # Default Colors
  primary:
    background: '0x212121'
    foreground: '0xf1f1f1'
  # Normal colors
  normal:
    black:   '0x212121'
    red:     '0xc30771'
    green:   '0x10a778'
    yellow:  '0xa89c14'
    blue:    '0x008ec4'
    magenta: '0x523c79'
    cyan:    '0x20a5ba'
    white:   '0xe0e0e0'
  # Bright colors
  bright:
    black:   '0x818181'
    red:     '0xfb007a'
    green:   '0x5fd7af'
    yellow:  '0xf3e430'
    blue:    '0x20bbfc'
    magenta: '0x6855de'
    cyan:    '0x4fb8cc'
    white:   '0xf1f1f1'
`

const PENCIL_LIGHT = `
# Colors (Pencil Light)
colors:
  # Default Colors
  primary:
    background: '0xf1f1f1'
    foreground: '0x424242'
  # Normal colors
  normal:
    black:   '0x212121'
    red:     '0xc30771'
    green:   '0x10a778'
    yellow:  '0xa89c14'
    blue:    '0x008ec4'
    magenta: '0x523c79'
    cyan:    '0x20a5ba'
    white:   '0xe0e0e0'
  # Bright colors
  bright:
    black:   '0x212121'
    red:     '0xfb007a'
    green:   '0x5fd7af'
    yellow:  '0xf3e430'
    blue:    '0x20bbfc'
    magenta: '0x6855de'
    cyan:    '0x4fb8cc'
    white:   '0xf1f1f1'
`

const REMEDY_DARK = `
colors:
  # Default colors
  primary:
    background: '0x2c2b2a'
    foreground: '0xf9e7c4'

    dim_foreground:    '0x685E4A'
    bright_foreground: '0x1C1508'
    dim_background:    '0x202322'
    bright_background: '0x353433'

  # Cursor colors
  cursor:
    text:   '0xf9e7c4'
    cursor: '0xf9e7c4'

  # Normal colors
  normal:
    black:   '0x282a2e'
    blue:    '0x5f819d'
    cyan:    '0x5e8d87'
    green:   '0x8c9440'
    magenta: '0x85678f'
    orange:  '0xcc6953'
    red:     '0xa54242'
    white:   '0x707880'
    yellow:  '0xde935f'

  # Bright colors
  bright:
    black:   '0x373b41'
    blue:    '0x81a2be'
    cyan:    '0x8abeb7'
    green:   '0xb5bd68'
    magenta: '0xb294bb'
    red:     '0xcc6666'
    white:   '0xc5c8c6'
    yellow:  '0xf0c674'
`

const SNAZZY = `
# Colors (Snazzy)
colors:
  # Default colors
  primary:
    background: '0x282a36'
    foreground: '0xeff0eb'

  # Normal colors
  normal:
    black:   '0x282a36'
    red:     '0xff5c57'
    green:   '0x5af78e'
    yellow:  '0xf3f99d'
    blue:    '0x57c7ff'
    magenta: '0xff6ac1'
    cyan:    '0x9aedfe'
    white:   '0xf1f1f0'

  # Bright colors
  bright:
    black:   '0x686868'
    red:     '0xff5c57'
    green:   '0x5af78e'
    yellow:  '0xf3f99d'
    blue:    '0x57c7ff'
    magenta: '0xff6ac1'
    cyan:    '0x9aedfe'
    white:   '0xf1f1f0'
`

const SOLARIZED_DARK = `
# Colors (Solarized Dark)
colors:
  # Default colors
  primary:
    background: '0x002b36'
    foreground: '0x839496'

  # Normal colors
  normal:
    black:   '0x073642'
    red:     '0xdc322f'
    green:   '0x859900'
    yellow:  '0xb58900'
    blue:    '0x268bd2'
    magenta: '0xd33682'
    cyan:    '0x2aa198'
    white:   '0xeee8d5'

  # Bright colors
  bright:
    black:   '0x002b36'
    red:     '0xcb4b16'
    green:   '0x586e75'
    yellow:  '0x657b83'
    blue:    '0x839496'
    magenta: '0x6c71c4'
    cyan:    '0x93a1a1'
    white:   '0xfdf6e3'
`

const SOLARIZED_LIGHT = `
# Colors (Solarized Light)
colors:
  # Default colors
  primary:
    background: '0xfdf6e3'
    foreground: '0x586e75'

  # Normal colors
  normal:
    black:   '0x073642'
    red:     '0xdc322f'
    green:   '0x859900'
    yellow:  '0xb58900'
    blue:    '0x268bd2'
    magenta: '0xd33682'
    cyan:    '0x2aa198'
    white:   '0xeee8d5'

  # Bright colors
  bright:
    black:   '0x002b36'
    red:     '0xcb4b16'
    green:   '0x586e75'
    yellow:  '0x657b83'
    blue:    '0x839496'
    magenta: '0x6c71c4'
    cyan:    '0x93a1a1'
    white:   '0xfdf6e3'
`

const STORM = `
404: Not Found
`

const TAERMINAL = `
# Colors (Taerminal)
colors:
  # Default colors
  primary:
    background: '0x26282a'
    foreground: '0xf0f0f0'
  cursor:
    background: '0xf0f0f0'
    foreground: '0x26282a'

  # Normal colors
  normal:
    black:   '0x26282a'
    red:     '0xff8878'
    green:   '0xb4fb73'
    yellow:  '0xfffcb7'
    blue:    '0x8bbce5'
    magenta: '0xffb2fe'
    cyan:    '0xa2e1f8'
    white:   '0xf1f1f1'

  # Bright colors
  bright:
    black:   '0x6f6f6f'
    red:     '0xfe978b'
    green:   '0xd6fcba'
    yellow:  '0xfffed5'
    blue:    '0xc2e3ff'
    magenta: '0xffc6ff'
    cyan:    '0xc0e9f8'
    white:   '0xffffff'
`

const TANGO_DARK = `
# GNOME Terminal Tango Dark
colors:
  primary:
    background: '0x2e3436'
    foreground: '0xd3d7cf'
    
  normal:
    black: '0x2e3436'
    red: '0xcc0000'
    green: '0x4e9a06'
    yellow: '0xc4a000'
    blue: '0x3465a4'
    magenta: '0x75507b'
    cyan: '0x06989a'
    white: '0xd3d7cf'

  bright:
    black: '0x555753'
    red: '0xef2929'
    green: '0x8ae234'
    yellow: '0xfce94f'
    blue: '0x729fcf'
    magenta: '0xad7fa8'
    cyan: '0x34e2e2'
    white: '0xeeeeec'
`

const TENDER = `
colors:
  # Default colors
  primary:
    background: '0x282828'
    foreground: '0xeeeeee'

  # Normal colors
  normal:
    black:   '0x282828'
    red:     '0xf43753'
    green:   '0xc9d05c'
    yellow:  '0xffc24b'
    blue:    '0xb3deef'
    magenta: '0xd3b987'
    cyan:    '0x73cef4'
    white:   '0xeeeeee'

  # Bright colors
  bright:
    black:   '0x4c4c4c'
    red:     '0xf43753'
    green:   '0xc9d05c'
    yellow:  '0xffc24b'
    blue:    '0xb3deef'
    magenta: '0xd3b987'
    cyan:    '0x73cef4'
    white:   '0xfeffff'
`

const TERMINAL_APP = `
# Colors (Terminal.app)
colors:
  # Default colors
  primary:
    background: '0x000000'
    foreground: '0xb6b6b6'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0x990000'
    green:   '0x00a600'
    yellow:  '0x999900'
    blue:    '0x0000b2'
    magenta: '0xb200b2'
    cyan:    '0x00a6b2'
    white:   '0xbfbfbf'

  # Bright colors
  bright:
    black:   '0x666666'
    red:     '0xe50000'
    green:   '0x00d900'
    yellow:  '0xe5e500'
    blue:    '0x0000ff'
    magenta: '0xe500e5'
    cyan:    '0x00e5e5'
    white:   '0xe5e5e5'
`

const THE_LOVELACE = `
colors:
  # Default colors
  primary:
    background: '0x1D1F28'
    foreground: '0xFDFDFD'

  # Normal colors
  normal:
  # Bright colors
    black:   '0x282A36'
    red:     '0xF37F97'
    green:   '0x5ADECD'
    yellow:  '0xF2A272'
    blue:    '0x8897F4'
    magenta: '0xC574DD'
    cyan:    '0x79E6F3'
    white:   '0xFDFDFD'
  bright:
    black:   '0x414458'
    red:     '0xFF4971'
    green:   '0x18E3C8'
    yellow:  '0xEBCB8B'
    blue:    '0xFF8037'
    magenta: '0x556FFF'
    cyan:    '0x3FDCEE'
    white:   '0xBEBEC1'
  indexed_colors: []
`

const TOMORROW_NIGHT = `
# Colors (Tomorrow Night)
colors:
  # Default colors
  primary:
    background: '0x1d1f21'
    foreground: '0xc5c8c6'

  cursor:
    text: '0x1d1f21'
    cursor: '0xffffff'

  # Normal colors
  normal:
    black:   '0x1d1f21'
    red:     '0xcc6666'
    green:   '0xb5bd68'
    yellow:  '0xe6c547'
    blue:    '0x81a2be'
    magenta: '0xb294bb'
    cyan:    '0x70c0ba'
    white:   '0x373b41'

  # Bright colors
  bright:
    black:   '0x666666'
    red:     '0xff3334'
    green:   '0x9ec400'
    yellow:  '0xf0c674'
    blue:    '0x81a2be'
    magenta: '0xb77ee0'
    cyan:    '0x54ced6'
    white:   '0x282a2e'
`

const TOMORROW_NIGHT_BRIGHT = `
# Colors (Tomorrow Night Bright)
colors:
  # Default colors
  primary:
    background: '0x000000'
    foreground: '0xeaeaea'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xd54e53'
    green:   '0xb9ca4a'
    yellow:  '0xe6c547'
    blue:    '0x7aa6da'
    magenta: '0xc397d8'
    cyan:    '0x70c0ba'
    white:   '0x424242'

  # Bright colors
  bright:
    black:   '0x666666'
    red:     '0xff3334'
    green:   '0x9ec400'
    yellow:  '0xe7c547'
    blue:    '0x7aa6da'
    magenta: '0xb77ee0'
    cyan:    '0x54ced6'
    white:   '0x2a2a2a'
`

const WOMBAT = `
# Colors (Wombat)
colors:
  # Default colors
  primary:
    background: '0x1f1f1f'
    foreground: '0xe5e1d8'

  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xf7786d'
    green:   '0xbde97c'
    yellow:  '0xefdfac'
    blue:    '0x6ebaf8'
    magenta: '0xef88ff'
    cyan:    '0x90fdf8'
    white:   '0xe5e1d8'

  # Bright colors
  bright:
    black:   '0xb4b4b4'
    red:     '0xf99f92'
    green:   '0xe3f7a1'
    yellow:  '0xf2e9bf'
    blue:    '0xb3d2ff'
    magenta: '0xe5bdff'
    cyan:    '0xc2fefa'
    white:   '0xffffff'
`

const XTERM = `
# XTerm's default colors
colors:
  # Default colors
  primary:
    background: '0x000000'
    foreground: '0xffffff'
  # Normal colors
  normal:
    black:   '0x000000'
    red:     '0xcd0000'
    green:   '0x00cd00'
    yellow:  '0xcdcd00'
    blue:    '0x0000ee'
    magenta: '0xcd00cd'
    cyan:    '0x00cdcd'
    white:   '0xe5e5e5'

  # Bright colors
  bright:
    black:   '0x7f7f7f'
    red:     '0xff0000'
    green:   '0x00ff00'
    yellow:  '0xffff00'
    blue:    '0x5c5cff'
    magenta: '0xff00ff'
    cyan:    '0x00ffff'
    white:   '0xffffff'
`
