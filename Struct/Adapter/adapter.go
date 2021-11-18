package adapter

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

func (adapter *EatAdapter) New(eat *Eat) {
	adapter.name = "eat adapter"
	adapter.adaptee = eat
}

func (adapter *EatAdapter) Do() string {
	return adapter.adaptee.Do()
}
