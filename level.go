package logger

// Level is used to determine what tag log messages should be given and whether they should be printed
type Level int

// Log levels used by the functions defined in logfunctions.go
const (
	// LogNone is solely used by SetLevel, to turn off all logging
	LogNone Level = iota
	// LogPanic is used by Recover and RecoverStack
	LogPanic
	// LogError is used by Error and Errorf
	LogError
	// LogWarn is used by Warn and Warnf
	LogWarn
	// LogNotice is used by Notice and Noticef
	LogNotice
	// LogInfo is used by Info and Infof
	LogInfo
	// LogDebug is used by Debug and Debugf
	LogDebug
	// LogTrace is used by Trace and Tracef
	LogTrace
	// LogAll is solely used by SetLevel, to enable all possible logging
	LogAll
)

var levelNames = map[Level]string{
	LogPanic:  "PANIC ",
	LogError:  "ERROR ",
	LogWarn:   "WARN  ",
	LogNotice: "NOTICE",
	LogInfo:   "INFO  ",
	LogDebug:  "DEBUG ",
	LogTrace:  "TRACE ",
}

// LevelName is a convenience method to convert a level to a string representation,
// the string returned is always 6 characters long
// For example: LevelName(logInfo) returns "INFO  "
func LevelName(level Level) string {
	return levelNames[level]
}
