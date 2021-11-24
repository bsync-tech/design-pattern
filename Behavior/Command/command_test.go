package command

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestGetInstance(c *C) {
	d1 := &GetRequest{Method: "GET"}
	d2 := &PostRequest{Method: "POST"}
	r1 := &Request{Request: d1}
	r2 := &Request{Request: d2}

	c.Assert(Action(r1) == "get method executed", Equals, true)
	c.Assert(Action(r2) == "post method executed", Equals, true)
}
