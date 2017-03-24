package logger

import "fmt"

type defaultLogger struct{}

func (defaultLogger) Log(level Level, msg string, stack []string) {
	fmt.Printf("[%s] %s\n", LevelName(level), msg)

	if stack != nil {
		for i, line := range stack {
			fmt.Printf("[%s] %d: %s\n", LevelName(level), len(stack)-i-1, line)
		}
	}
}
