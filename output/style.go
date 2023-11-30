package output

import (
	"fmt"
	"strconv"
	"strings"
)

type ColorSize uint8
type Typo uint8
type Text string
type StyleOption func(o *Style)

const (
	Size4Bit ColorSize = iota
	Size8Bit
	Size24Bit

	Foreground Typo = iota
	Background

	EOL = "\n"
	ESC = "\033["

	Black         uint64 = 30
	Red                  = 31
	Green                = 32
	Yellow               = 33
	Blue                 = 34
	Magenta              = 35
	Cyan                 = 36
	White                = 37
	BrightBlack          = 90
	BrightRed            = 91
	BrightGreen          = 92
	BrightYellow         = 93
	BrightBlue           = 94
	BrightMagenta        = 95
	BrightCyan           = 96
	BrightWhite          = 97
	Reset                = 0
)

var colors16 []uint64 = []uint64{Black, Red, Green, Yellow, Blue, Magenta, Cyan, White}
var styleMap = map[string]*Style{}
var creset = Color{Size4Bit, 0, 0, nil}

type Color struct {
	size  ColorSize
	color uint64
	typo  Typo
	err   error
}

func New16Color(c uint64, t Typo) Color {
	if contains(colors16, c) {
		return Color{Size4Bit, ternaryIf(t == Background, c+10, c), t, nil}
	}

	return Color{Size4Bit, 0, 0, fmt.Errorf("Color %d no es un color válido de 4 bits", c)}
}

func New256Color(c uint64, t Typo) Color {
	if c < 256 {
		return Color{Size8Bit, c, t, nil}
	}

	return Color{Size8Bit, 0, 0, fmt.Errorf("Color %d no es un color válido de 8 bits", c)}
}

func NewRGBColor(r, g, b uint16, t Typo) Color {
	if r < 256 && g < 256 && b < 256 {
		return Color{Size24Bit, uint64(r)<<16 | uint64(g)<<8 | uint64(b), t, nil}
	}

	return Color{Size24Bit, 0, 0, fmt.Errorf("Color rgb(%d,%d,%d) no es un color RGB válido de 24 bits", r, g, b)}
}

func NewHexColor(hex string, t Typo) Color {
	c := Color{}
	c.color, c.err = hex2Color(hex)
	c.typo = t
	c.size = Size24Bit

	return c
}

func (c Color) Error() error {
	return c.err
}

func (c Color) IsValid() bool {
	return c.err == nil
}

func (c Color) String() string {
	switch c.size {
	case Size4Bit:
		return fmt.Sprintf("%s%dm", ESC, c.color)
	case Size8Bit:
		return fmt.Sprintf("%s%d;5;%dm", ESC, ternaryIf(c.typo == Foreground, 38, 48), c.color)
	case Size24Bit:
		r := uint8(c.color >> 16)
		g := uint8((c.color >> 8) & 0xFF)
		b := uint8(c.color & 0xFF)
		return fmt.Sprintf("%s%d;2;%d;%d;%dm", ESC, ternaryIf(c.typo == Foreground, 38, 48), r, g, b)
	}

	return ""
}

type Style struct {
	foreground Color
	background Color
	dim        Text
	italic     Text
	underline  Text
	blinking   Text
	reverse    Text
	hidden     Text
	strike     Text
}

func contains[T comparable](s []T, f T) bool {
	for _, v := range s {
		if v == f {
			return true
		}
	}

	return false
}

func hex2Color(hex string) (c uint64, err error) {
	hex = strings.TrimPrefix(hex, "#")
	c, err = strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		err = fmt.Errorf("Color %s no es un color hexadecimal válido de 24 bits", hex)
	}

	return
}

func (s Style) String() string {
	return s.foreground.String() + s.background.String() + string(s.dim) + string(s.italic) + string(s.underline) +
		string(s.blinking) + string(s.reverse) + string(s.hidden) + string(s.strike)
}

func init() {
	reset := NewStyle(WithForeground(creset))
	AddStyle("reset", reset)
	AddStyle("/", reset)
}

// NewStyle crea una nueva instancia de Style
func NewStyle(opts ...StyleOption) *Style {
	s := &Style{}
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithForeground asigna foreground a Style
func WithForeground(foreground Color) StyleOption {
	return func(s *Style) {
		s.foreground = foreground
	}
}

// WithBackground asigna background a Style
func WithBackground(background Color) StyleOption {
	return func(s *Style) {
		s.background = background
	}
}

// WithDim asigna dim a Style
func WithDim() StyleOption {
	return func(s *Style) {
		s.dim = Text(fmt.Sprintf("%s%sm", ESC, Dim))
	}
}

// WithItalic asigna italic a Style
func WithItalic() StyleOption {
	return func(s *Style) {
		s.italic = Text(fmt.Sprintf("%s%sm", ESC, Italic))
	}
}

// WithUnderline asigna underline a Style
func WithUnderline() StyleOption {
	return func(s *Style) {
		s.underline = Text(fmt.Sprintf("%s%sm", ESC, Underline))
	}
}

// WithBlinking asigna blinking a Style
func WithBlinking() StyleOption {
	return func(s *Style) {
		s.blinking = Text(fmt.Sprintf("%s%sm", ESC, Blink))
	}
}

// WithReverse asigna reverse a Style
func WithReverse() StyleOption {
	return func(s *Style) {
		s.reverse = Text(fmt.Sprintf("%s%sm", ESC, Reverse))
	}
}

// WithHidden asigna hidden a Style
func WithHidden() StyleOption {
	return func(s *Style) {
		s.hidden = Text(fmt.Sprintf("%s%sm", ESC, Hidden))
	}
}

// WithStrike asigna strike a Style
func WithStrike() StyleOption {
	return func(s *Style) {
		s.strike = Text(fmt.Sprintf("%s%sm", ESC, Strike))
	}
}

func AddStyle(key string, style *Style) {
	styleMap[key] = style
}
