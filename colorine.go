// Package colorine provides a simple feature to print messages to console.
//
// A log line consists of a prefix and a message. The prefix is colored by its
// content (for example below, "create" is green, "exists" is blue, and "error"
// is red, and so on.) and the message is not colored.
//	$ your_program
//	    create path/to/file
//	    exists path/to/another/file
//	     error something went wrong
//
// By default, no prefix color is defined. You must specify prefix-to-color
// mapping when you create a Logger. See the Logger's example.
package colorine

import (
	"fmt"

	"github.com/daviddengcn/go-colortext"
)

// A Logger handles the printing stuff. Create one with NewLogger().
type Logger struct {
	Prefixes     map[string]TextStyle
	DefaultStyle TextStyle
}

// A TextStyle represents the style that strings are displayed with. Consists
// of a foreground and a background, each of which is one of the text color constants.
type TextStyle struct {
	Foreground textColor
	Background textColor
}

type textColor struct {
	Color  ct.Color
	Bright bool
}

// Text color constants. Use these to build a TextStyle.
var (
	None = textColor{ct.None, false}

	Black   = textColor{ct.Black, false}
	Red     = textColor{ct.Red, false}
	Green   = textColor{ct.Green, false}
	Yellow  = textColor{ct.Yellow, false}
	Blue    = textColor{ct.Blue, false}
	Magenta = textColor{ct.Magenta, false}
	Cyan    = textColor{ct.Cyan, false}
	White   = textColor{ct.White, false}

	BrightBlack   = textColor{ct.Black, true}
	BrightRed     = textColor{ct.Red, true}
	BrightGreen   = textColor{ct.Green, true}
	BrightYellow  = textColor{ct.Yellow, true}
	BrightBlue    = textColor{ct.Blue, true}
	BrightMagenta = textColor{ct.Magenta, true}
	BrightCyan    = textColor{ct.Cyan, true}
	BrightWhite   = textColor{ct.White, true}
)

// Predefined TextStyles. Of course you can ignore these.
var (
	Verbose = TextStyle{White, None}
	Info    = TextStyle{Green, None}
	Notice  = TextStyle{Blue, None}
	Warn    = TextStyle{Yellow, None}
	Error   = TextStyle{Red, None}
)

// NewLogger creates a new Logger.  prefixes is a prefix-string-to-text-style
// table. The log prefix is colored using this table. If no entry is found in
// prefixes, defaultStyle is used.
//
// By default, no prefix is registered to any text style.
func NewLogger(prefixes map[string]TextStyle, defaultStyle TextStyle) *Logger {
	return &Logger{prefixes, defaultStyle}
}

// Log prints messages to console. prefix is colored using the logger.prefixes
// table.
func (logger *Logger) Log(prefix, message string) {
	textColor, ok := logger.Prefixes[prefix]
	if !ok {
		textColor = logger.DefaultStyle
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
