package zap_test

import (
	"github.com/goph/logur"

	zapintegration "logur.dev/integration/zap"
)

func ExampleNew() {
	logger := zapintegration.New(logur.NewNoopLogger())

	// Output:
	_ = logger
}
