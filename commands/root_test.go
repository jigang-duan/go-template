package commands

import (
	qt "github.com/frankban/quicktest"
	"testing"
)

func TestRootCmdExecute(t *testing.T) {
	c := qt.New(t)

	resp := Execute([]string{})

	c.Assert(resp.Err, qt.IsNil)
}
