package proxy

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
	d1 := &TeslaCar{Name: "I"}
	c.Assert(d1.Drive() == "I drive tesla", Equals, true)
	d2 := &TeslaCarProxy{Name: "designated driver", ICar: *d1}
	fmt.Println(d2.Drive())
	c.Assert(d2.Drive() == "designated driver work for I drive tesla", Equals, true)
}
