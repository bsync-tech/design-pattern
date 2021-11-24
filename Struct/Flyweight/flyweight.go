package flyweight

import (
	"fmt"
)

type FlyWeight struct {
	CompanyName    string
	CompanyAddress string
	CompanyWebsite string
}

func (f FlyWeight) String() string {
	return fmt.Sprintf("%s,%s,%s", f.CompanyName, f.CompanyAddress, f.CompanyWebsite)
}

type Employee struct {
	ShareInfo *FlyWeight
	Name      string
}

func (e *Employee) String() string {
	return fmt.Sprintf("%s,%s", e.ShareInfo, e.Name)
}
