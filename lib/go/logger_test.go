package logger_test

import (
	"testing"

	logger "github.com/takashabe/bazel-exercise/lib/go/"
)

func TestFoo(t *testing.T) {
	logger.Info("hoge")
}
