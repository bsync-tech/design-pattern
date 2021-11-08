package adapter

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &Eat{name: "eat"}
	c.Assert(d1.Do() == "eat", Equals, true)
	d2 := EatAdapter{adaptee: d1}
	c.Assert(d2.Do() == "eat adapter", Equals, true)
}
