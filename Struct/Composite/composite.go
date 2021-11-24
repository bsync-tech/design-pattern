package composite

import (
	"fmt"
	"strings"
)

type ObjectInterface interface {
	String() string
}

type Object struct {
	name string
}

func (s Object) String() string {
	return s.name
}

type CompositeObject struct {
	objects []ObjectInterface
}

func (s *CompositeObject) AddObjects(objects ...ObjectInterface) {
	s.objects = append(s.objects, objects...)
}

func (s CompositeObject) String() string {
	var str []string
	for _, v := range s.objects {
		str = append(str, fmt.Sprintf("%s", v))
	}
	return strings.Join(str, ",")
}
