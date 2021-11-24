package proxy

import (
	"fmt"
)

type CarInterface interface {
	Drive() string
}

type TeslaCar struct {
	Name string
}

func (t TeslaCar) Drive() string {
	return fmt.Sprintf("%s drive tesla", t.Name)
}

type TeslaCarProxy struct {
	ICar CarInterface
	Name string
}

func (t TeslaCarProxy) Drive() string {
	return fmt.Sprintf("%s work for %s", t.Name, t.ICar.Drive())
}
