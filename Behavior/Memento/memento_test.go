package memento

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	a1 := &Memento{}

	c1 := &MemoManager{}
	c1.SetMemento(a1)

	d1 := &Originator{}
	d1.SetState("n1")
	c1.GetMemento().SetSnapshot(d1.GetState())
	c.Assert(c1.GetMemento().GetSnapshot() == "n1", Equals, true)
	d1.SetState("n2")
	c.Assert(c1.GetMemento().GetSnapshot() != d1.GetState(), Equals, true)
	d1.SetState(c1.GetMemento().GetSnapshot())
	c.Assert(c1.GetMemento().GetSnapshot() == "n1", Equals, true)
}
