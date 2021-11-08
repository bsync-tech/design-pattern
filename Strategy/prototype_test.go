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
	d1 := GetInstance()
	d2 := GetInstance()
	c.Assert(d1 == d2, Equals, true)
	c.Assert(d1.GetName() == d2.GetName(), Equals, true)
}
