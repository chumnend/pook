package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// start test runner
	code := m.Run()
	os.Exit(code)
}
