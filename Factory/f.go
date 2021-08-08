package f

// In this pattern, Add Duck type need, you need
// 1. Add lines code in FactoryProduceDuck
// 2. Add new type Duck in this file

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

type DuckFactory interface {
	ProduceDuck() Duck
}

type BeijingDuckFactory struct {
	name string
}

func (s BeijingDuckFactory) ProduceDuck() Duck {
	return BeijingDuck{
		name: "BeijingDuck",
	}
}

type ShaoxingDuckFactory struct {
	name string
}

func (s ShaoxingDuckFactory) ProduceDuck() Duck {
	return ShaoxingDuck{
		name: "ShaoxingDuck",
	}
}
