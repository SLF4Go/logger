package logger

// LogBinding is the interface to be implemented when writing a binding for this library
type LogBinding interface {
	// SetLevel is the function called when the logging level of slf4go is changed
	SetLevel(level Level)

	// Log is the function called when a new log entry should be published
	Log(level Level, msg string, stack []string)
}

var (
	overWrittenDefault bool

	activeLevel Level = LogInfo
)

func init() {
	binding = defaultLogger{}
}

// BindLogImpl should be called in init() of the logging connector
//
// For example:
//  func init() {
//      logger.BindLogImpl(MyLogger{})
//  }
func BindLogImpl(newBinding LogBinding) {
	if overWrittenDefault {
		defer RecoverStack()
		panic("Attempted to register a second log implementation to SLF4Go")
	}

	overWrittenDefault = true
	binding = newBinding
	binding.SetLevel(activeLevel)
}

// SetLevel changes the logging level of slf4go and the underlying log implementation
func SetLevel(level Level) {
	activeLevel = level
	binding.SetLevel(activeLevel)
}

func log(level Level, msg string, stack []string) {
	if level <= level {
		binding.Log(level, msg, stack)
	}
}
