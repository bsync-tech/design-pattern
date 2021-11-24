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
	bf := new(BeijingDuckFactory)
	c.Assert(bf == nil, Equals, false)
	c.Logf("BeijingDuckFactory produce %s color, %s shape", bf.ProduceColor().GetName(), bf.ProduceShape().GetName())
	c.Assert(bf.ProduceColor().GetName() == "BlueDuck", Equals, true)
	c.Assert(bf.ProduceShape().GetName() == "BigDuck", Equals, true)
	sf := new(ShaoxingDuckFactory)
	c.Assert(sf == nil, Equals, false)
	c.Logf("ShaoxingDuckFactory produce %s color, %s shape", sf.ProduceColor().GetName(), sf.ProduceShape().GetName())
	c.Assert(sf.ProduceColor().GetName() == "YelloDuck", Equals, true)
	c.Assert(sf.ProduceShape().GetName() == "SmallDuck", Equals, true)
	fmt.Println(c.GetTestLog())
}
