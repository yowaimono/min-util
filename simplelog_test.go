package minutil

import (
	"testing"
)

func TestSimpleLog(t *testing.T) {
	Info("This is an info message")
	Error("This is an error message")
	Warn("This is a warning message")
	Debug("This is a debug message")
}