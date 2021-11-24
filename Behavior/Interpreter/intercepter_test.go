package intercepter

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &Context{}
	d1.Init()
	r1 := d1.Interpret("Beijing Women")
	c.Assert(r1 == "Free ride", Equals, true)
	r2 := d1.Interpret("Beijing Man")
	c.Assert(r2 == "Pay please", Equals, true)
}
