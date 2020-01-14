package loggers

import (
	qt "github.com/frankban/quicktest"
	"testing"
)

func TestLoger(t *testing.T) {
	c := qt.New(t)
	l := NewWarningLogger()

	l.ERROR.Println("One error")
	l.ERROR.Println("Two error")
	l.WARN.Println("A warning")

	c.Assert(l.ErrorCounter.Count(), qt.Equals, uint64(2))
}
