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
	bfd := bf.ProduceDuck("BeijingDuck")
	c.Assert(bf == nil, Equals, false)
	c.Logf("BeijingDuckFactory produce 'BeijingDuck', name is %s", bfd.GetName())
	c.Assert(bfd.GetName() == "BeijingDuck", Equals, true)
	sfd := bf.ProduceDuck("ShaoxingDuck")
	c.Logf("BeijingDuckFactory produce 'ShaoxingDuck', name is %s", sfd.GetName())
	c.Assert(sfd.GetName() == "FakeShaoxingDuck", Equals, true)

	sf := new(ShaoxingDuckFactory)
	sfd = sf.ProduceDuck("ShaoxingDuck")
	c.Assert(sfd == nil, Equals, false)
	c.Logf("ShaoxingDuckFactory produce 'ShaoxingDuck', name is %s", sfd.GetName())
	c.Assert(sfd.GetName() == "ShaoxingDuck", Equals, true)
	bfd = sf.ProduceDuck("BeijingDuck")
	c.Logf("ShaoxingDuckFactory produce 'BeijingDuck', name is %s", bfd.GetName())
	c.Assert(bfd.GetName() == "FakeBeijingDuck", Equals, true)

	fmt.Println(c.GetTestLog())
}
