package composite

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &Object{name: "d1"}
	c.Assert(d1.String() == "d1", Equals, true)
	d2 := &Object{name: "d2"}
	c.Assert(d2.String() == "d2", Equals, true)
	d3 := &CompositeObject{}
	d3.AddObjects(*d1, *d2)
	c.Assert(d3.String() == "d1,d2", Equals, true)
}
