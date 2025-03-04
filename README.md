# Go logger

## Installation

```bash
go get package github.com/mrborghini/go-logger
```

## Usage

```go
package main

import (
	"github.com/mrborghini/go-logger"
)

func main() {
	logger := logger.NewLogger("main")
	logger.Info("Hello, World!")
	logger.Warning("Hello warning!")
	logger.Error("Hello, Error!")
	logger.Debug("Hello, Debug!")
}
```

If you want to enable the debug log level then you need to add:

`DEBUG=true` to your environment variables.
