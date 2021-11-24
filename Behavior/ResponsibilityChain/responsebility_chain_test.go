package rchain

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &HandlerA{Name: "d1"}
	d2 := &HandlerB{Name: "d2"}
	d1.SetNext(d2)

	c.Assert(d1.HandleRequest("need A") == "request need A processed by d1", Equals, true)
	c.Assert(d1.HandleRequest("need B") == "request need B processed by d2", Equals, true)
}
