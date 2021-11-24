package iterator

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &ConcretContainer{}
	d1.Objs = append(d1.Objs, "1")
	d1.Objs = append(d1.Objs, "2")
	d1.Objs = append(d1.Objs, "3")

	var result []string
	it := d1.GetIterator()
	for {
		if it.HasNext() == false {
			break
		}
		obj := it.Next()
		result = append(result, obj.(string))
	}
	c.Assert(strings.Join(result, ",") == "1,2,3", Equals, true)
}
