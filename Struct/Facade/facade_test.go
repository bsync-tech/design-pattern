package facade

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &ComputerFacade{Cpu: new(Cpu), Memory: new(Memory), Disk: new(Disk)}
	c.Assert(d1.Run() == "cpu load instructions,memory load data,disk read,cpu execute", Equals, true)
}
