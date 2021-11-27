package adapter

import "fmt"

type ICharge interface {
	Supply() string
}

type Charge220 struct {
	name string
}

func (c *Charge220) Charge() string {
	return c.name
}

type ChargeAdapter struct {
	name    string
	adaptee *Charge220
}

func (a *ChargeAdapter) New(e *Charge220) *ChargeAdapter {
	a.name = "'s adapter"
	a.adaptee = e
	return a
}

func (a *ChargeAdapter) Charge() string {
	return fmt.Sprintf("%s%s", a.adaptee.Charge(), a.name)
}
