package colorine

import "testing"

func TestLog(t *testing.T) {
	logger := NewLogger(
		map[string]TextColor{
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
