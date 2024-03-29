package testlog

import (
	"bytes"
	"io"
	"log"
	"sync"
	"testing"
)

// Catch redirects the given loggers output to a bytes.Buffer with
// cleared flags. Restoring them once the test is done. If no loggers
// are given log.Default() is used.
func Catch(t *testing.T, loggers ...Logger) *Trap {
	if len(loggers) == 0 {
		loggers = append(loggers, log.Default())
	}
	var trap Trap
	for _, l := range loggers {
		w, f := l.Writer(), l.Flags()
		reset := func() {
			l.SetOutput(w)
			l.SetFlags(f)
		}
		t.Cleanup(reset)

		// use buffer
		l.SetOutput(&trap)
		l.SetFlags(0)
	}
	return &trap
}

// Wrap returns a test func which replaces the default log output and
// uses t.Log if the test fails.
func Wrap(test testFunc) testFunc {
	return func(t *testing.T) {
		t.Helper()
		trap := Catch(t)
		test(t)
		if t.Failed() && trap.Len() > 0 {
			t.Log("Wrapped log output\n", trap.String())
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

type Trap struct {
	m   sync.RWMutex
	buf bytes.Buffer
}

func (l *Trap) Write(p []byte) (n int, err error) {
	l.m.Lock()
	n, err = l.buf.Write(p)
	l.m.Unlock()
	return
}

func (l *Trap) String() string {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.buf.String()
}

func (l *Trap) Len() int {
	return l.buf.Len()
}
