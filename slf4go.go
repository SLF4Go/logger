package logger

// The interface to be implemented when writing a connector for this library
type LogImpl interface {
	Log(level Level, msg string, stack []string)
}

var overWrittenDefault bool

// This method should be called in init() of the logging connector
//
// For example:
//  func init() {
//      logger.RegisterLogImpl(MyLogger{})
//  }
func RegisterLogImpl(newImpl LogImpl) {
	if overWrittenDefault {
		defer RecoverStack()
		panic("Attempted to register a second log implementation to SLF4Go")
		return
	}

	overWrittenDefault = true
	impl = newImpl
}
