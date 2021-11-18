package builder

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	director := new(Director)
	builder := new(HighPerformanceComputerBuilder)
	director.SetBuilder(builder)
	computer := director.Build()
	c.Assert(computer.String() == "amd 3995wx ps-2 logitech", Equals, true)
}
