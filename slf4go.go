package logger

import (
	"fmt"
)

type LogImpl interface {
	Log(level Level, msg string, stack []string)
}

var implementations = make(map[string]*LogImpl)

func RegisterLogImpl(name string, impl LogImpl) {
	if len(implementations) == 0 {
		defaultCtx.impl = impl
	}
	implementations[name] = &impl
}

func New(implName string) *LogContext {
	impl, exists := implementations[implName]
	if !exists {
		panic(fmt.Sprintf("logger impl with name '%s' does not exist", implName))
	}
	return &LogContext{*impl}
}
