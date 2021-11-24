package state

type IState interface {
	Action(*Context) string
}

type StateInit struct{}

func (s *StateInit) Action(c *Context) string {
	c.State = s
	return "StateInit"
}

type StateA struct{}

func (s *StateA) Action(c *Context) string {
	c.State = s
	return "StateA"
}

type StateB struct{}

func (s *StateB) Action(c *Context) string {
	c.State = s
	return "StateB"
}

type StateStop struct{}

func (s *StateStop) Action(c *Context) string {
	c.State = s
	return "StateStop"
}

type Context struct {
	State IState
}

func (c *Context) Init() {
	c.State = new(StateInit)
}
