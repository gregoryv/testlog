package testlog_test

import (
	"log"
	"testing"
	"time"

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

// run with go test -race -count 100 to see if it ever fails
func TestCatch_safe(t *testing.T) {
	buf := testlog.Catch(t)
	go log.Print("x") // a write
	<-time.After(time.Millisecond)
	if got := buf.String(); got != "x\n" { // and a read
		t.Error(got)
	}
}

func TestTrap_Len(t *testing.T) {
	var trap testlog.Trap
	trap.Write([]byte("123"))
	if got := trap.Len(); got != 3 {
		t.Error(got)
	}
}
