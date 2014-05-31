package colorine

import "testing"

func TestLog(t *testing.T) {
	logger := NewLogger(
		map[string]TextStyle{
			"verbose": Verbose,
			"info":    Info,
			"notice":  Notice,
			"warn":    Warn,
			"error":   Error,
		},
		Info,
	)

	logger.Log("verbose", "shows this color")
	logger.Log("info", "shows this color")
	logger.Log("notice", "shows this color")
	logger.Log("warn", "shows this color")
	logger.Log("error", "shows this color")
}

func ExampleLogger() {
	logger := NewLogger(
		map[string]TextStyle{
			// Using colorine's predefined TextStyles
			"create": Info,
			"exists": Notice,
			// Or specify your own
			"debug": TextStyle{White, None},
			"error": TextStyle{BrightRed, White},
		},
		// The default prefix color
		TextStyle{Green, None},
	)

	logger.Log("create", "path/to/file")
	logger.Log("exists", "path/to/another/file")
	logger.Log("error", "something went wrong")
}
