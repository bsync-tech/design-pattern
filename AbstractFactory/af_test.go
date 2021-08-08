package af

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestProduceDuck(c *C) {
	BeijingDuck := FactoryProduceDuck("BeijingDuck")
	c.Assert(BeijingDuck == nil, Equals, false)
	c.Logf("BeijingDuck's name is %s", BeijingDuck.GetName())
	ShaoxingDuck := FactoryProduceDuck("ShaoxingDuck")
	c.Assert(ShaoxingDuck == nil, Equals, false)
	c.Assert(ShaoxingDuck.GetName(), Equals, "ShaoxingDuck")
	c.Logf("ShaoxingDuck's name is %s", ShaoxingDuck.GetName())
	OtherDuck := FactoryProduceDuck("OtherDuck")
	c.Assert(OtherDuck == nil, Equals, true)
	fmt.Println(c.GetTestLog())
}
