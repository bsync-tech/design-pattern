package observer

import (
	"fmt"
	"strings"
)

type ISubject interface {
	Add(ob IObserver)
	Remove(ob IObserver)
	NotifyObservers()
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Add(ob IObserver) {
	if s.observers == nil {
		s.observers = make([]IObserver, 0)
	}
	s.observers = append(s.observers, ob)
}
func (s *Subject) Remove(ob IObserver) {
	for i, v := range s.observers {
		if ob == v {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}
func (s *Subject) NotifyObserver() string {
	var resp []string
	for _, v := range s.observers {
		r := v.Response()
		resp = append(resp, r)
	}
	return strings.Join(resp, ",")
}

type IObserver interface {
	Response() string
}

type Observer struct {
	Name string
}

func (o *Observer) Response() string {
	return fmt.Sprintf("%s response", o.Name)
}
