package decorator

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &SmallDuck{Name: "small duck"}
	c.Assert(d1.String() == "small duck", Equals, true)
	d2 := &SmallDuckDecorator{Duck: *d1}
	d2.WithHat()
	c.Assert(d2.String() == "small duck,with red hat", Equals, true)
}
