package logger

type LogImpl interface {
	Log(level Level, msg string, stack []string)
}

var overWrittenDefault bool

func RegisterLogImpl(newImpl LogImpl) {
	if overWrittenDefault {
		defer RecoverStack()
		panic("Attempted to register a second log implementation to SLF4Go")
		return
	}

	overWrittenDefault = true
	impl = newImpl
}
