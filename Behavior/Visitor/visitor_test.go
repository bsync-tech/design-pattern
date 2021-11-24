package visitor

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	o1 := &Object{
		ElementA: &ElementA{Value: "ElementA"},
		ElementB: &ElementB{Value: "ElementB"},
	}
	ElementGetValueAvisitor := &GetValueVisitor{}
	c.Assert(o1.Accept(ElementGetValueAvisitor) == "ElementA,ElementB", Equals, true)
}
