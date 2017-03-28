package logger

import (
	"fmt"
	"runtime/debug"
	"strings"
	"unicode/utf8"
)

var impl LogImpl

// Calls the underlying log implementation with log level LogError and the provided message
func Error(msg string) {
	impl.Log(LogError, msg, nil)
}

// Calls the underlying log implementation with log level LogError and a formatted message
func Errorf(msg string, args ...interface{}) {
	Error(fmt.Sprintf(msg, args...))
}

// Calls the underlying log implementation with log level LogWarn and the provided message
func Warn(msg string) {
	impl.Log(LogWarn, msg, nil)
}

// Calls the underlying log implementation with log level LogWarn and a formatted message
func Warnf(msg string, args ...interface{}) {
	Warn(fmt.Sprintf(msg, args...))
}

// Calls the underlying log implementation with log level LogNotice and the provided message
func Notice(msg string) {
	impl.Log(LogNotice, msg, nil)
}

// Calls the underlying log implementation with log level LogNotice and a formatted message
func Noticef(msg string, args ...interface{}) {
	Notice(fmt.Sprintf(msg, args...))
}

// Calls the underlying log implementation with log level LogInfo and the provided message
func Info(msg string) {
	impl.Log(LogInfo, msg, nil)
}

// Calls the underlying log implementation with log level LogInfo and a formatted message
func Infof(msg string, args ...interface{}) {
	Info(fmt.Sprintf(msg, args...))
}

// Calls the underlying log implementation with log level LogDebug and the provided message
func Debug(msg string) {
	impl.Log(LogDebug, msg, nil)
}

// Calls the underlying log implementation with log level LogDebug and a formatted message
func Debugf(msg string, args ...interface{}) {
	Debug(fmt.Sprintf(msg, args...))
}

// Calls the underlying log implementation with log level LogTrace and the provided message
func Trace(msg string) {
	impl.Log(LogTrace, msg, nil)
}

// Calls the underlying log implementation with log level LogTrace and a formatted message
func Tracef(msg string, args ...interface{}) {
	Trace(fmt.Sprintf(msg, args...))
}

// Alternative method to recover() that invokes the underlying log implementation
// with log level LogPanic and attempts to format the panic message with %v
//
// Usage:
//  defer logger.Recover()
func Recover() {
	if panicMsg := recover(); panicMsg != nil {
		recoverInternal(panicMsg, nil)
	}
}

// Alternative method to recover() that invokes the underlying log implementation
// with log level LogPanic and attempts to format the panic message with %v.
// Next to that the logger will also output the stack trace that led up to the panic()
func RecoverStack() {
	if panicMsg := recover(); panicMsg != nil {
		recoverInternal(panicMsg, debug.Stack())
	}
}

func recoverInternal(panic interface{}, stack []byte) {
	description := fmt.Sprintf("%v", panic)
	if stack == nil {
		impl.Log(LogPanic, description, nil)
	} else {
		stackTrace := string(stack)
		stackLines := strings.Split(stackTrace, "\n")

		var stack = make([]string, 0, len(stackLines)-6)

		for i, line := range stackLines {
			if i > 6 && utf8.RuneCountInString(line) > 0 {
				stack = append(stack, line)
			}
		}

		impl.Log(LogPanic, description, stack)
	}
}
