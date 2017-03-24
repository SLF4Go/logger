package logger

import (
	"fmt"
	"runtime/debug"
	"strings"
	"unicode/utf8"
)

type LogContext struct {
	impl LogImpl
}

func (ctx *LogContext) Error(msg string) {
	ctx.impl.Log(LogError, msg, nil)
}

func (ctx *LogContext) ErrorF(msg string, args ...interface{}) {
	ctx.Error(fmt.Sprintf(msg, args...))
}

func (ctx *LogContext) Warn(msg string) {
	ctx.impl.Log(LogWarn, msg, nil)
}

func (ctx *LogContext) WarnF(msg string, args ...interface{}) {
	ctx.Warn(fmt.Sprintf(msg, args...))
}

func (ctx *LogContext) Notice(msg string) {
	ctx.impl.Log(LogNotice, msg, nil)
}

func (ctx *LogContext) NoticeF(msg string, args ...interface{}) {
	ctx.Notice(fmt.Sprintf(msg, args...))
}

func (ctx *LogContext) Info(msg string) {
	ctx.impl.Log(LogInfo, msg, nil)
}

func (ctx *LogContext) InfoF(msg string, args ...interface{}) {
	ctx.Info(fmt.Sprintf(msg, args...))
}

func (ctx *LogContext) Debug(msg string) {
	ctx.impl.Log(LogDebug, msg, nil)
}

func (ctx *LogContext) DebugF(msg string, args ...interface{}) {
	ctx.Debug(fmt.Sprintf(msg, args...))
}

func (ctx *LogContext) Trace(msg string) {
	ctx.impl.Log(LogTrace, msg, nil)
}

func (ctx *LogContext) TraceF(msg string, args ...interface{}) {
	ctx.Trace(fmt.Sprintf(msg, args...))
}

func (ctx *LogContext) Recover() {
	if panicMsg := recover(); panicMsg != nil {
		ctx.recoverInternal(panicMsg, nil)
	}
}

func (ctx *LogContext) RecoverStack() {
	if panicMsg := recover(); panicMsg != nil {
		ctx.recoverInternal(panicMsg, debug.Stack())
	}
}

func (ctx *LogContext) recoverInternal(panic interface{}, stack []byte) {
	description := fmt.Sprintf("Message: %v", panic)
	if stack == nil {
		ctx.impl.Log(LogPanic, description, nil)
	} else {
		stackTrace := string(stack)
		stackLines := strings.Split(stackTrace, "\n")

		var stack = make([]string, 0, len(stackLines)-6)

		for i, line := range stackLines {
			if i > 6 && utf8.RuneCountInString(line) > 0 {
				stack = append(stack, line)
			}
		}

		ctx.impl.Log(LogPanic, description, stack)
	}
}
