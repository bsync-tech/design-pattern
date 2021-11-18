package flyweight

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &FlyWeight{CompanyName: "company name", CompanyAddress: "company address", CompanyWebsite: "company website"}
	d2 := &Employee{ShareInfo: d1, Name: "d2 name"}
	d3 := &Employee{ShareInfo: d1, Name: "d3 name"}
	d4 := &Employee{ShareInfo: d1, Name: "d4 name"}
	c.Assert(d2.ShareInfo == d3.ShareInfo, Equals, true)
	c.Assert(d3.ShareInfo == d4.ShareInfo, Equals, true)
}
