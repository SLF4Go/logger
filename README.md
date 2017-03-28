# SLF4Go
[![Build Status](https://travis-ci.org/slf4go/logger.svg?branch=master)](https://travis-ci.org/slf4go/logger)
[![GoDoc](https://godoc.org/github.com/slf4go/logger?status.svg)](https://godoc.org/github.com/slf4go/logger)

SLF4go, or in full 'Simple Logging Facade 4 Go' is a logging interface much like [SLF4J](https://www.slf4j.org/). This library allows APIs/Libraries to call upon a generic logging interface while still leaving the choice of actual logging library to the end developer.

## Using SLF4Go in your project
##### For library developers
```go
import "github.com/slf4go/logger"
```

#### For end developers
This is an example on how to choose a specific logging library. In this example I'm showing [go-logging](https://github.com/op/go-logging)
```go
import (
	"github.com/slf4go/logger"
	_ "github.com/slf4go/go-logging-connector"
)
```

After having done this, you can still configure your specific logging module (go-logging here) to your hearts content. Any libraries you use that use SLF4Go will automatically use this logging library as well
