package logger

import "fmt"

// This type is used to determine what tag log messages should be given and whether they should be printed
type Level int

// Log levels used by the functions defined in logfunctions.go
const (
	// Log level used by Recover and RecoverStack
	LogPanic Level = iota
	// Log level used by Error and Errorf
	LogError
	// Log level used by Warn and Warnf
	LogWarn
	// Log level used by Notice and Noticef
	LogNotice
	// Log level used by Info and Infof
	LogInfo
	// Log level used by Debug and Debugf
	LogDebug
	// Log level used by Trace and Tracef
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

// Convenience method to convert a level to a string representation
// For example: LevelName(logInfo) returns "INFO"
func LevelName(level Level) string {
	return fmt.Sprintf("%-6s", levelNames[level])
}
