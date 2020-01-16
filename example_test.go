package zap_test

import (
	"logur.dev/logur"

	zapintegration "logur.dev/integration/zap"
)

func ExampleNew() {
	logger := zapintegration.New(logur.NoopLogger{})

	// Output:
	_ = logger
}
