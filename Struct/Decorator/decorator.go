package decorator

import "fmt"

type Duck interface {
	String() string
}

type SmallDuck struct {
	Name string
}

func (s SmallDuck) String() string {
	return s.Name
}

type SmallDuckDecoratorInterface interface {
	WithHat()
}

type SmallDuckDecorator struct {
	Duck Duck
	Hat  string
}

func (d *SmallDuckDecorator) WithHat() {
	d.Hat = "with red hat"
}
func (d SmallDuckDecorator) String() string {
	return fmt.Sprintf("%s,%s", d.Duck, d.Hat)
}
