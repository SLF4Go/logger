package logger

import (
	"fmt"
	"runtime/debug"
	"strings"
	"unicode/utf8"
)

var impl LogImpl

func Error(msg string) {
	impl.Log(LogError, msg, nil)
}

func Errorf(msg string, args ...interface{}) {
	Error(fmt.Sprintf(msg, args...))
}

func Warn(msg string) {
	impl.Log(LogWarn, msg, nil)
}

func Warnf(msg string, args ...interface{}) {
	Warn(fmt.Sprintf(msg, args...))
}

func Notice(msg string) {
	impl.Log(LogNotice, msg, nil)
}

func Noticef(msg string, args ...interface{}) {
	Notice(fmt.Sprintf(msg, args...))
}

func Info(msg string) {
	impl.Log(LogInfo, msg, nil)
}

func Infof(msg string, args ...interface{}) {
	Info(fmt.Sprintf(msg, args...))
}

func Debug(msg string) {
	impl.Log(LogDebug, msg, nil)
}

func Debugf(msg string, args ...interface{}) {
	Debug(fmt.Sprintf(msg, args...))
}

func Trace(msg string) {
	impl.Log(LogTrace, msg, nil)
}

func Tracef(msg string, args ...interface{}) {
	Trace(fmt.Sprintf(msg, args...))
}

func Recover() {
	if panicMsg := recover(); panicMsg != nil {
		recoverInternal(panicMsg, nil)
	}
}

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
