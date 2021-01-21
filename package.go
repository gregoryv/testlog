package testlog

import (
	"bytes"
	"log"
	"testing"
)

// Wrap returns a test func which replaces the default log output and
// uses t.Log if the test fails.
func Wrap(test testFunc) testFunc {
	return func(t *testing.T) {
		t.Helper()
		var buf bytes.Buffer
		log.SetOutput(&buf)
		test(t)
		if t.Failed() && buf.Len() > 0 {
			t.Log("Wrapped log output\n", buf.String())
		}
	}
}

type testFunc func(*testing.T)
