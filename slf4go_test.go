package logger

import (
	"fmt"
	"testing"
)

// Create a special test log implementation that validates values passed
type testLogImpl struct {
	def      defaultLogger
	validate func(level Level, msg string, stack []string)
}

func (impl testLogImpl) Log(level Level, msg string, stack []string) {
	impl.def.Log(level, msg, stack)
	impl.validate(level, msg, stack)
}

var (
	// Context that we shall be calling
	testCtx *LogContext
	// Keep this reference so we can adjust the validator
	testImpl *testLogImpl
)

// Create a reference so we can adjust the test validator and then register it
func init() {
	testImpl = &testLogImpl{def: defaultLogger{}}
	RegisterLogImpl("test", testImpl)
}

func TestNew(t *testing.T) {
	testCtx = New("test")
}

func TestError(t *testing.T) {
	// Set up test
	const testMsg = "This is an ERROR message!"
	testImpl.validate = func(level Level, msg string, stack []string) {
		if level != LogError || msg != testMsg || stack != nil {
			t.FailNow()
		}
	}

	// Run test
	testCtx.Error(testMsg)
}

func TestWarn(t *testing.T) {
	// Set up test
	const testMsg = "This is a WARN message!"
	testImpl.validate = func(level Level, msg string, stack []string) {
		if level != LogWarn || msg != testMsg || stack != nil {
			t.FailNow()
		}
	}

	// Run test
	testCtx.Warn(testMsg)
}

func TestInfo(t *testing.T) {
	// Set up test
	const testMsg = "This is an INFO message!"
	testImpl.validate = func(level Level, msg string, stack []string) {
		if level != LogInfo || msg != testMsg || stack != nil {
			t.FailNow()
		}
	}

	// Run test
	testCtx.Info(testMsg)
}

func TestDebug(t *testing.T) {
	// Set up test
	const testMsg = "This is a DEBUG message!"
	testImpl.validate = func(level Level, msg string, stack []string) {
		if level != LogDebug || msg != testMsg || stack != nil {
			t.FailNow()
		}
	}

	// Run test
	testCtx.Debug(testMsg)
}

func TestTrace(t *testing.T) {
	// Set up test
	const testMsg = "This is a TRACE message!"
	testImpl.validate = func(level Level, msg string, stack []string) {
		if level != LogTrace || msg != testMsg || stack != nil {
			t.FailNow()
		}
	}

	// Run test
	testCtx.Trace(testMsg)
}

func TestRecover(t *testing.T) {
	// Set up test
	const panicMsg = "This panic shouldn't have a stack trace."
	testImpl.validate = func(level Level, msg string, stack []string) {
		compareMsg := fmt.Sprintf("Message: %s", panicMsg)
		if level != LogPanic || msg != compareMsg || stack != nil {
			t.FailNow()
		}
	}

	// Run test
	defer testCtx.Recover()
	panic(panicMsg)

	// Test should never reach this point
	t.FailNow()
}

func TestRecoverStack(t *testing.T) {
	// Set up test
	const panicMsg = "This panic should have a stack trace."
	testImpl.validate = func(level Level, msg string, stack []string) {
		compareMsg := fmt.Sprintf("Message: %s", panicMsg)
		if level != LogPanic || msg != compareMsg || stack == nil {
			t.FailNow()
		}
	}

	// Run test
	defer testCtx.RecoverStack()
	panic(panicMsg)

	// Test should never reach this point
	t.FailNow()
}
