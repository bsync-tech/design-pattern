package prototype

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := new(Duck)
	d2 := d1.Clone()
	c.Assert(*d1 == d2, Equals, true)
	c.Assert(d1.String() == d2.String(), Equals, true)
}
