package adapter

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &Eat{name: "eat"}
	c.Assert(d1.Do() == "eat", Equals, true)
	d2 := new(EatAdapter).New(d1)

	fmt.Println("hi", d2.Do())
	c.Assert(d2.Do() == "eat's adapter", Equals, true)
}
