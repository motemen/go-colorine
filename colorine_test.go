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
	// Output:
	//[0;37m   verbose[0m shows this color
	//[0;32m      info[0m shows this color
	//[0;34m    notice[0m shows this color
	//[0;33m      warn[0m shows this color
	//[0;31m     error[0m shows this color
}
