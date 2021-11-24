package f

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
	bfd := bf.ProduceDuck()
	c.Assert(bf == nil, Equals, false)
	c.Logf("BeijingDuckFactory produce 'BeijingDuck', name is %s", bfd.GetName())
	c.Assert(bfd.GetName() == "BeijingDuck", Equals, true)

	sf := new(ShaoxingDuckFactory)
	sfd := sf.ProduceDuck()
	c.Assert(sfd == nil, Equals, false)
	c.Logf("ShaoxingDuckFactory produce 'ShaoxingDuck', name is %s", sfd.GetName())
	c.Assert(sfd.GetName() == "ShaoxingDuck", Equals, true)

	fmt.Println(c.GetTestLog())
}
