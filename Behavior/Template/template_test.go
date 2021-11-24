package template

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	t := &ISortTemplate{}
	d1 := &SortA{}
	t.SetSort(d1)
	c.Assert(t.Sort() == "SortA init,SortA shuffle,SortA merge", Equals, true)
	d2 := &SortB{}
	t.SetSort(d2)
	c.Assert(t.Sort() == "SortB init,SortB shuffle,SortB merge", Equals, true)
}
