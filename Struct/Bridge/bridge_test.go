package bridge

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := new(Device)
	d1.SetBrand(&AppleBrand{})
	d1.SetDeviceType(&PhoneDeviceType{})
	c.Assert(d1.String() == "apple-phone", Equals, true)
	d2 := new(Device)
	d2.SetBrand(&HuaweiBrand{})
	d2.SetDeviceType(&PadDeviceType{})
	c.Assert(d2.String() == "huawei-pad", Equals, true)
}
