package logger

import (
	"fmt"
	"runtime/debug"
	"strings"
	"unicode/utf8"
)

var binding LogBinding

// Error calls the underlying log implementation with log level LogError and the provided message
func Error(msg string) {
	log(LogError, msg, nil)
}

// ErrorE calls the underlying log implementation with log level LogError and the result of err.Error()
func ErrorE(err error) {
	log(LogError, err.Error(), nil)
}

// Errorf calls the underlying log implementation with log level LogError and a formatted message
func Errorf(format string, args ...interface{}) {
	Error(fmt.Sprintf(format, args...))
}

// Warn calls the underlying log implementation with log level LogWarn and the provided message
func Warn(msg string) {
	log(LogWarn, msg, nil)
}

// Warnf calls the underlying log implementation with log level LogWarn and a formatted message
func Warnf(format string, args ...interface{}) {
	Warn(fmt.Sprintf(format, args...))
}

// Notice calls the underlying log implementation with log level LogNotice and the provided message
func Notice(msg string) {
	log(LogNotice, msg, nil)
}

// Noticef calls the underlying log implementation with log level LogNotice and a formatted message
func Noticef(format string, args ...interface{}) {
	Notice(fmt.Sprintf(format, args...))
}

// Info calls the underlying log implementation with log level LogInfo and the provided message
func Info(msg string) {
	log(LogInfo, msg, nil)
}

// Infof calls the underlying log implementation with log level LogInfo and a formatted message
func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}

// Debug calls the underlying log implementation with log level LogDebug and the provided message
func Debug(msg string) {
	log(LogDebug, msg, nil)
}

// Debugf calls the underlying log implementation with log level LogDebug and a formatted message
func Debugf(format string, args ...interface{}) {
	Debug(fmt.Sprintf(format, args...))
}

// Trace calls the underlying log implementation with log level LogTrace and the provided message
func Trace(msg string) {
	log(LogTrace, msg, nil)
}

// Tracef calls the underlying log implementation with log level LogTrace and a formatted message
func Tracef(format string, args ...interface{}) {
	Trace(fmt.Sprintf(format, args...))
}

// Recover is an alternative method to recover() that invokes the underlying log implementation
// with log level LogPanic and attempts to format the panic message with %v
//
// Usage:
//  defer logger.Recover()
func Recover() {
	if panicMsg := recover(); panicMsg != nil {
		recoverInternal(panicMsg, nil)
	}
}

// RecoverStack is an alternative method to recover() that invokes the underlying log implementation
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
		log(LogPanic, description, nil)
	} else {
		stackTrace := string(stack)
		stackLines := strings.Split(stackTrace, "\n")

		var stack = make([]string, 0, len(stackLines)-6)

		for i, line := range stackLines {
			if i > 6 && utf8.RuneCountInString(line) > 0 {
				stack = append(stack, line)
			}
		}

		log(LogPanic, description, stack)
	}
}
