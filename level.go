package logger

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Level int

const (
	LogPanic Level = iota
	LogError
	LogWarn
	LogInfo
	LogDebug
	LogTrace
)

var nameLength int
var levelNames map[Level]string

func init() {
	levelNames = make(map[Level]string)

	RegisterCustomLevel(LogPanic, "PANIC")
	RegisterCustomLevel(LogError, "ERROR")
	RegisterCustomLevel(LogWarn, "WARN")
	RegisterCustomLevel(LogInfo, "INFO")
	RegisterCustomLevel(LogDebug, "DEBUG")
	RegisterCustomLevel(LogTrace, "TRACE")
}

func LevelName(level Level) string {
	return fmt.Sprintf("%-"+strconv.Itoa(nameLength)+"s", levelNames[level])
}

func RegisterCustomLevel(level Level, name string) {
	if length := utf8.RuneCountInString(name); length > nameLength {
		nameLength = length
	}

	levelNames[level] = name
}
