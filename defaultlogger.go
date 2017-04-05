package logger

import (
	"fmt"
	"os"
)

var gaveWarning bool

type defaultLogger struct{}

func (l defaultLogger) Log(level Level, msg string, stack []string) {
	if !gaveWarning {
		gaveWarning = true
		l.Log(LogWarn, "You are using the SLF4Go default logger, it is recommended to import a connector", nil)
	}

	var out *os.File
	if level <= LogWarn {
		out = os.Stderr
	} else {
		out = os.Stdout
	}

	fmt.Fprintf(out, "[%s] %s\n", LevelName(level), msg)

	if stack != nil {
		for i, line := range stack {
			fmt.Fprintf(out, "[%s] %d: %s\n", LevelName(level), len(stack)-i-1, line)
		}
	}
}

func (defaultLogger) SetLevel(level Level) {}
