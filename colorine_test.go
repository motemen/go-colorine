package colorine_test

import "github.com/motemen/go-colorine"
import "testing"

func TestLog(t *testing.T) {
	logger := colorine.NewLogger(
		colorine.Prefixes{
			"verbose": colorine.Verbose,
			"info":    colorine.Info,
			"notice":  colorine.Notice,
			"warn":    colorine.Warn,
			"error":   colorine.Error,
		},
		colorine.Info,
	)

	logger.Log("verbose", "shows this color")
	logger.Log("info", "shows this color")
	logger.Log("notice", "shows this color")
	logger.Log("warn", "shows this color")
	logger.Log("error", "shows this color")
}

func ExampleLogger() {
	logger := colorine.NewLogger(
		colorine.Prefixes{
			// Using colorine's predefined TextStyles
			"create": colorine.Info,
			"exist":  colorine.Notice,
			// Or specify your own
			"debug": colorine.TextStyle{colorine.White, colorine.None},
			"error": colorine.TextStyle{colorine.BrightRed, colorine.White},
		},
		// The default prefix color
		colorine.TextStyle{colorine.Green, colorine.None},
	)

	logger.Log("create", "path/to/file")
	logger.Log("exist", "path/to/another/file")
	logger.Log("error", "something went wrong!")
}
