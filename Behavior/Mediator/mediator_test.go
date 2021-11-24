package mediator

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &Mediator{}
	d2 := &Employee{Mediator: d1}
	d3 := &Employer{Mediator: d1}
	c.Assert(d2.Hunt() == "employee send to employer with need a job", Equals, true)
	c.Assert(d3.Seek() == "employer send to employee with hunt one", Equals, true)
}
