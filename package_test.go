package testlog_test

import (
	"log"
	"testing"

	"github.com/gregoryv/testlog"
)

func TestWrap(t *testing.T) {
	t.Run("x", testlog.Wrap(func(t *testing.T) {
		log.Print("test")
		// uncomment to see results
		//t.Fail()
	}))
}

func TestCatch(t *testing.T) {
	buf := testlog.Catch(t)
	log.Print("x")
	if got := buf.String(); got != "x\n" {
		t.Error(got)
	}
}
