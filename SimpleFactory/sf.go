package sf

// In this pattern, Add Duck type need, you need
// 1. Add lines code in FactoryProduceDuck
// 2. Add new type Duck in this file

type Duck interface {
	GetName() string
}

type DuckFactory struct{}

func (s *DuckFactory) ProduceDuck(typ string) Duck {
	switch typ {
	case "BeijingDuck":
		return BeijingDuck{
			name: "BeijingDuck",
		}
	case "ShaoxingDuck":
		return ShaoxingDuck{
			name: "ShaoxingDuck",
		}
	}
	return nil
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
