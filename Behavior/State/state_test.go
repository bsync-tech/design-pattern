package state

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	ctx := &Context{}
	ctx.Init()

	s1 := new(StateA)
	c.Assert(s1.Action(ctx) == "StateA", Equals, true)

	s2 := new(StateB)
	c.Assert(s2.Action(ctx) == "StateB", Equals, true)

	s3 := new(StateStop)
	c.Assert(s3.Action(ctx) == "StateStop", Equals, true)
}
