package colorine

import (
	"fmt"

	"github.com/daviddengcn/go-colortext"
)

// The logger. Create one with NewLogger().
type Logger struct {
	Prefixes     map[string]TextStyle
	DefaultStyle TextStyle
}

type TextStyle struct {
	Foreground TextColor
	Background TextColor
}

type TextColor struct {
	Color  ct.Color
	Bright bool
}

// Text color constants.
var (
	None = TextColor{ct.None, false}

	Black   = TextColor{ct.Black, false}
	Red     = TextColor{ct.Red, false}
	Green   = TextColor{ct.Green, false}
	Yellow  = TextColor{ct.Yellow, false}
	Blue    = TextColor{ct.Blue, false}
	Magenta = TextColor{ct.Magenta, false}
	Cyan    = TextColor{ct.Cyan, false}
	White   = TextColor{ct.White, false}

	BrightBlack   = TextColor{ct.Black, true}
	BrightRed     = TextColor{ct.Red, true}
	BrightGreen   = TextColor{ct.Green, true}
	BrightYellow  = TextColor{ct.Yellow, true}
	BrightBlue    = TextColor{ct.Blue, true}
	BrightMagenta = TextColor{ct.Magenta, true}
	BrightCyan    = TextColor{ct.Cyan, true}
	BrightWhite   = TextColor{ct.White, true}
)

// Predefined text styles. Of course you can ignore these.
var (
	Verbose = TextStyle{White, None}
	Info    = TextStyle{Green, None}
	Notice  = TextStyle{Blue, None}
	Warn    = TextStyle{Yellow, None}
	Error   = TextStyle{Red, None}
)

// Creates a new Logger.
// prefixes is a prefix-string-to-text-style table. The log prefix is colored
// using this table. If no entry is found in prefixes, defaultStyle is used.
//
// By default, no prefix is registered to any text style.
func NewLogger(prefixes map[string]TextStyle, defaultStyle TextStyle) *Logger {
	return &Logger{prefixes, defaultStyle}
}

// Log messages to console. prefix is colored using the logger.prefixes table.
func (logger *Logger) Log(prefix, message string) {
	TextColor, ok := logger.Prefixes[prefix]
	if !ok {
		TextColor = logger.DefaultStyle
	}

	ct.ChangeColor(
		TextColor.Foreground.Color,
		TextColor.Foreground.Bright,
		TextColor.Background.Color,
		TextColor.Background.Bright,
	)

	fmt.Printf("%10s", prefix)

	ct.ResetColor()

	fmt.Printf(" %s\n", message)
}
