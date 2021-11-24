package mediator

import "fmt"

type IMediator interface {
	Nofity(from string, to string, msg string) string
}

type Mediator struct{}

func (c *Mediator) Nofity(from string, to string, msg string) string {
	return fmt.Sprintf("%s send to %s with %s", from, to, msg)
}

type Employer struct {
	Mediator IMediator
}

func (e *Employer) Seek() string {
	return e.Mediator.Nofity("employer", "employee", "hunt one")
}

type Employee struct {
	Mediator IMediator
}

func (e *Employee) Hunt() string {
	return e.Mediator.Nofity("employee", "employer", "need a job")
}
