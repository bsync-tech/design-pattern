package f

// In this pattern, Add Duck type, you need:
// 1. Add new type Duck in this file.
// 2. Add new Factory type in this file.

type Duck interface {
	GetName() string
}

type BeijingDuck struct {
	name string
}

func (b BeijingDuck) GetName() string {
	return b.name
}

type ShaoxingDuck struct {
	name string
}

func (s ShaoxingDuck) GetName() string {
	return s.name
}

type IDuckFactory interface {
	ProduceDuck() Duck
}

type BeijingDuckFactory struct{}

func (s BeijingDuckFactory) ProduceDuck() Duck {
	return BeijingDuck{
		name: "BeijingDuck",
	}
}

type ShaoxingDuckFactory struct{}

func (s ShaoxingDuckFactory) ProduceDuck() Duck {
	return ShaoxingDuck{
		name: "ShaoxingDuck",
	}
}
