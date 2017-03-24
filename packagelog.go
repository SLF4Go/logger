package logger

import (
	"runtime/debug"
)

var defaultCtx LogContext

func init() {
	defaultCtx = LogContext{defaultLogger{}}
}

func Error(msg string) {
	defaultCtx.Error(msg)
}

func ErrorF(msg string, args ...interface{}) {
	defaultCtx.ErrorF(msg, args...)
}

func Warn(msg string) {
	defaultCtx.Warn(msg)
}

func WarnF(msg string, args ...interface{}) {
	defaultCtx.WarnF(msg, args...)
}

func Notice(msg string) {
	defaultCtx.Notice(msg)
}

func NoticeF(msg string, args ...interface{}) {
	defaultCtx.NoticeF(msg, args...)
}

func Info(msg string) {
	defaultCtx.Info(msg)
}

func InfoF(msg string, args ...interface{}) {
	defaultCtx.InfoF(msg, args...)
}

func Debug(msg string) {
	defaultCtx.Debug(msg)
}

func DebugF(msg string, args ...interface{}) {
	defaultCtx.DebugF(msg, args...)
}

func Trace(msg string) {
	defaultCtx.Trace(msg)
}

func TraceF(msg string, args ...interface{}) {
	defaultCtx.TraceF(msg, args...)
}

func Recover() {
	if panicMsg := recover(); panicMsg != nil {
		defaultCtx.recoverInternal(panicMsg, nil)
	}
}

func RecoverStack() {
	if panicMsg := recover(); panicMsg != nil {
		defaultCtx.recoverInternal(panicMsg, debug.Stack())
	}
}
