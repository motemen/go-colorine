package colorine

import (
	"fmt"

	"github.com/daviddengcn/go-colortext"
)

type Color struct {
	Color  ct.Color
	Bright bool
}

type TextColor struct {
	Foreground Color
	Background Color
}

var (
	None = Color{ct.None, false}

	Black   = Color{ct.Black, false}
	Red     = Color{ct.Red, false}
	Green   = Color{ct.Green, false}
	Yellow  = Color{ct.Yellow, false}
	Blue    = Color{ct.Blue, false}
	Magenta = Color{ct.Magenta, false}
	Cyan    = Color{ct.Cyan, false}
	White   = Color{ct.White, false}

	BrightBlack   = Color{ct.Black, true}
	BrightRed     = Color{ct.Red, true}
	BrightGreen   = Color{ct.Green, true}
	BrightYellow  = Color{ct.Yellow, true}
	BrightBlue    = Color{ct.Blue, true}
	BrightMagenta = Color{ct.Magenta, true}
	BrightCyan    = Color{ct.Cyan, true}
	BrightWhite   = Color{ct.White, true}
)

var (
	Verbose = TextColor{White, None}
	Info    = TextColor{Green, None}
	Notice  = TextColor{Blue, None}
	Warn    = TextColor{Yellow, None}
	Error   = TextColor{Red, None}
)

type Logger struct {
	Prefixes         map[string]TextColor
	DefaultTextColor TextColor
}

func NewLogger(prefixes map[string]TextColor, defaultTextColor TextColor) *Logger {
	return &Logger{prefixes, defaultTextColor}
}

func (logger *Logger) Log(prefix, message string) {
	textColor, ok := logger.Prefixes[prefix]
	if !ok {
		textColor = logger.DefaultTextColor
	}

	ct.ChangeColor(
		textColor.Foreground.Color,
		textColor.Foreground.Bright,
		textColor.Background.Color,
		textColor.Background.Bright,
	)

	fmt.Printf("%10s", prefix)

	ct.ResetColor()

	fmt.Printf(" %s\n", message)
}
