package testlog_test

import (
	"log"
	"testing"

	"github.com/gregoryv/testlog"
)

func TestCatch(t *testing.T) {
	buf := testlog.Catch(t)
	log.Print("x")
	if got := buf.String(); got != "x\n" {
		t.Error(got)
	}
}
