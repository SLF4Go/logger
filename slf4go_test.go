package logger

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// Create a special test log implementation that validates values passed
type testLogImpl struct {
	validate func(level Level, msg string, stack []string)
}

func (impl testLogImpl) Log(level Level, msg string, stack []string) {
	impl.validate(level, msg, stack)
}

func (impl testLogImpl) SetLevel(level Level) {}

// We keep our own type-asserted pointer so we can set the validate func later on
var testImpl *testLogImpl

// Create a reference so we can adjust the test validator and then register it
func init() {
	testImpl = &testLogImpl{}
	BindLogImpl(testImpl)
}

func testLogFunction(t *testing.T, call func(string, ...interface{}), expectedLevel Level) {
	const defaultFormat string = "log test %d"
	random := rand.Intn(100)

	testImpl.validate = func(level Level, response string, stack []string) {
		expectedResponse := fmt.Sprintf(defaultFormat, random)
		if level != expectedLevel || response != expectedResponse || stack != nil {
			t.FailNow()
		}
	}

	call(defaultFormat, random)
}

func TestError(t *testing.T) {
	testLogFunction(t, Errorf, LogError)
}

func TestWarn(t *testing.T) {
	testLogFunction(t, Warnf, LogWarn)
}

func TestNotice(t *testing.T) {
	testLogFunction(t, Noticef, LogNotice)
}

func TestInfo(t *testing.T) {
	testLogFunction(t, Infof, LogInfo)
}

func TestDebug(t *testing.T) {
	testLogFunction(t, Debugf, LogDebug)
}

func TestTrace(t *testing.T) {
	testLogFunction(t, Tracef, LogTrace)
}

func TestRecover(t *testing.T) {
	defer Recover()
	testLogFunction(t, func(msg string, args ...interface{}) {
		panic(fmt.Sprintf(msg, args...))
	}, LogPanic)

	// Test should never reach this point
	t.FailNow()
}

func TestRecoverStack(t *testing.T) {
	defer RecoverStack()

	const defaultFormat string = "log test %d"
	random := rand.Intn(100)

	testImpl.validate = func(level Level, response string, stack []string) {
		expectedResponse := fmt.Sprintf(defaultFormat, random)
		if level != LogPanic || response != expectedResponse || stack == nil {
			t.FailNow()
		}

		if !strings.Contains(stack[0], "TestRecoverStack") {
			t.FailNow()
		}
	}

	panic(fmt.Sprintf(defaultFormat, random))

	// Test should never reach this point
	t.FailNow()
}
