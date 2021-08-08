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
	c.Logf("BeijingDuck's name is %s", bfd.GetName())
	sf := new(ShaoxingDuckFactory)
	sfd := sf.ProduceDuck()
	c.Assert(sfd == nil, Equals, false)
	c.Assert(sfd.GetName(), Equals, "ShaoxingDuck")
	c.Logf("ShaoxingDuck's name is %s", sfd.GetName())
	fmt.Println(c.GetTestLog())
}
