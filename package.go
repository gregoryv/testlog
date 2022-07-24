package testlog

import (
	"bytes"
	"io"
	"log"
	"testing"
)

// Catch redirects the given loggers output to a bytes.Buffer with
// cleared flags. Restoring them once the test is done. If no loggers
// are given log.Default() is used.
func Catch(t *testing.T, loggers ...Logger) *bytes.Buffer {
	if len(loggers) == 0 {
		loggers = append(loggers, log.Default())
	}
	var buf bytes.Buffer
	for _, l := range loggers {
		w, f := l.Writer(), l.Flags()
		reset := func() {
			l.SetOutput(w)
			l.SetFlags(f)
		}
		t.Cleanup(reset)

		// use buffer
		l.SetOutput(&buf)
		l.SetFlags(0)
	}
	return &buf
}

// Wrap returns a test func which replaces the default log output and
// uses t.Log if the test fails.
func Wrap(test testFunc) testFunc {
	return func(t *testing.T) {
		t.Helper()
		buf := Catch(t)
		test(t)
		if t.Failed() && buf.Len() > 0 {
			t.Log("Wrapped log output\n", buf.String())
		}
	}
}

type testFunc func(*testing.T)

type Logger interface {
	Writer() io.Writer
	Flags() int
	SetFlags(int)
	SetOutput(io.Writer)
}
