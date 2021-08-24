package af

// In this pattern, If need produce new Duck Product, you need
// 1. Add new type Duck in this file
// 2. Add new DuckFactory type

type Color interface {
	GetName() string
}

type YelloDuck struct {
	name string
}

func (y YelloDuck) GetName() string {
	return y.name
}

type BlueDuck struct {
	name string
}

func (b BlueDuck) GetName() string {
	return b.name
}

type Shape interface {
	GetName() string
}

type BigDuck struct {
	name string
}

func (b BigDuck) GetName() string {
	return b.name
}

type SmallDuck struct {
	name string
}

func (s SmallDuck) GetName() string {
	return s.name
}

type DuckFactory interface {
	ProduceColor() Color
	ProduceShape() Shape
}

type BeijingDuckFactory struct{}

func (b BeijingDuckFactory) ProduceColor() Color {
	return BlueDuck{
		name: "BlueDuck",
	}
}

func (b BeijingDuckFactory) ProduceShape() Shape {
	return BigDuck{
		name: "BigDuck",
	}
}

type ShaoxingDuckFactory struct{}

func (s ShaoxingDuckFactory) ProduceColor() Color {
	return BlueDuck{
		name: "YelloDuck",
	}
}

func (s ShaoxingDuckFactory) ProduceShape() Shape {
	return SmallDuck{
		name: "SmallDuck",
	}
}
