package observer

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &Subject{}
	d2 := &Observer{Name: "d2"}
	d3 := &Observer{Name: "d3"}
	d1.Add(d2)
	d1.Add(d3)
	c.Assert(d1.NotifyObserver() == "d2 response,d3 response", Equals, true)
	d1.Remove(d2)
	c.Assert(d1.NotifyObserver() == "d3 response", Equals, true)
}
