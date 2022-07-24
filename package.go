package testlog

import (
	"bytes"
	"log"
	"testing"
)

// Catch sets log.Default() writer to a bytes buffer and clears it's
// flags.  Restoring them once the test is done.
func Catch(t *testing.T) *bytes.Buffer {
	w := log.Default().Writer()
	f := log.Default().Flags()
	t.Cleanup(func() {
		log.SetOutput(w)
		log.SetFlags(f)
	})
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
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
