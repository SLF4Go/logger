package logger

import "fmt"

type Level int

const (
	LogPanic Level = iota
	LogError
	LogWarn
	LogNotice
	LogInfo
	LogDebug
	LogTrace
)

var levelNames = map[Level]string{
	LogPanic:  "PANIC",
	LogError:  "ERROR",
	LogWarn:   "WARN",
	LogNotice: "NOTICE",
	LogInfo:   "INFO",
	LogDebug:  "DEBUG",
	LogTrace:  "TRACE",
}

func LevelName(level Level) string {
	return fmt.Sprintf("%-6s", levelNames[level])
}
