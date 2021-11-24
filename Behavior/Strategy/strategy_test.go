package strategy

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	c1 := &Context{}
	s1 := &StrategyA{}
	c1.SetSrategy(s1)
	c.Assert(c1.Process() == "process with StrategyA", Equals, true)
	s2 := &StrategyB{}
	c1.SetSrategy(s2)
	c.Assert(c1.Process() == "process with StrategyB", Equals, true)
}
