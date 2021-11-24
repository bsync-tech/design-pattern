package intercepter

import (
	"strings"
)

type Expression interface {
	Interpret(string) bool
}
type TerminalExpression struct {
	nodes map[string]bool
}

func (t *TerminalExpression) SetNodes(nodes []string) {
	t.nodes = make(map[string]bool)
	for _, v := range nodes {
		t.nodes[v] = true
	}
}
func (t *TerminalExpression) Interpret(info string) bool {
	if _, ok := t.nodes[info]; ok {
		return true
	}
	return false
}

type AndExpression struct {
	city   Expression
	people Expression
}

func (and *AndExpression) Interpret(info string) bool {

	ss := strings.Split(info, " ")

	return and.city.Interpret(ss[0]) && and.people.Interpret(ss[1])
}

type Context struct {
	cities   []string
	people   []string
	freeRide Expression
}

func (cx *Context) Init() {
	cx.cities = []string{"Beijing", "Shanghai"}
	cx.people = []string{"Old", "Women", "Child"}
	city := new(TerminalExpression)
	people := new(TerminalExpression)
	city.SetNodes(cx.cities)
	people.SetNodes(cx.people)
	cx.freeRide = &AndExpression{city, people}
}

func (cx *Context) Interpret(info string) string {
	if ok := cx.freeRide.Interpret(info); ok {
		return "Free ride"
	} else {
		return "Pay please"
	}

}
