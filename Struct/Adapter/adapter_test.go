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
	d1 := &Charge220{name: "220v"}
	c.Assert(d1.Charge() == "220v", Equals, true)
	d2 := new(ChargeAdapter).New(d1)

	fmt.Println("hi", d2.Charge())
	c.Assert(d2.Charge() == "220v's adapter", Equals, true)
}
