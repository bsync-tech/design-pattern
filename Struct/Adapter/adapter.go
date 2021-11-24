package adapter

import "fmt"

type GeneralEat interface {
	Do() string
}

type Eat struct {
	name string
}

func (eat *Eat) New(name string) {
	eat.name = name
}

func (e *Eat) Do() string {
	return e.name
}

type EatAdapter struct {
	name    string
	adaptee *Eat
}

func (adapter *EatAdapter) New(eat *Eat) *EatAdapter {
	adapter.name = "'s adapter"
	adapter.adaptee = eat
	return adapter
}

func (adapter *EatAdapter) Do() string {
	return fmt.Sprintf("%s%s", adapter.adaptee.Do(), adapter.name)
}
