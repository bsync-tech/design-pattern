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

type FakeBeijingDuck struct {
	name string
}

func (b FakeBeijingDuck) GetName() string {
	return b.name
}

type ShaoxingDuck struct {
	name string
}

func (s ShaoxingDuck) GetName() string {
	return s.name
}

type FakeShaoxingDuck struct {
	name string
}

func (s FakeShaoxingDuck) GetName() string {
	return s.name
}

type DuckFactory interface {
	ProduceDuck() Duck
}

type BeijingDuckFactory struct{}

func (s BeijingDuckFactory) ProduceDuck(typ string) Duck {
	switch typ {
	case "BeijingDuck":
		return BeijingDuck{
			name: "BeijingDuck",
		}
	case "ShaoxingDuck":
		return FakeBeijingDuck{
			name: "FakeShaoxingDuck",
		}
	}
	return nil
}

type ShaoxingDuckFactory struct{}

func (s ShaoxingDuckFactory) ProduceDuck(typ string) Duck {
	switch typ {
	case "ShaoxingDuck":
		return ShaoxingDuck{
			name: "ShaoxingDuck",
		}
	case "BeijingDuck":
		return FakeBeijingDuck{
			name: "FakeBeijingDuck",
		}
	}
	return nil
}
