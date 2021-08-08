package af

// In this pattern, Add Duck type need, you need
// 1. Add lines code in FactoryProduceDuck
// 2. Add new type Duck in this file

type Duck interface {
	GetName() string
}

type CookDuckFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type DuckFactory struct{}

// FactoryProduceDuck
func FactoryProduceDuck(typ string) Duck {
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

type Cook interface {
	GetCookName() string
}

// FactoryProduceCook
func FactoryProduceCook(typ string) Cook {
	switch typ {
	case "SaltWaterCook":
		return SaltWaterCook{
			name: "SaltWaterCook",
		}
	case "RoastCook":
		return RoastCook{
			name: "RoastCook",
		}
	case "BeerCook":
		return BeerCook{
			name: "BeerCook",
		}
	}
	return nil
}

type SaltWaterCook struct {
	name string
}

func (s SaltWaterCook) GetCookName() string {
	return s.name
}

type RoastCook struct {
	name string
}

func (r RoastCook) GetCookName() string {
	return r.name
}

type BeerCook struct {
	name string
}

func (b BeerCook) GetCookName() string {
	return b.name
}
