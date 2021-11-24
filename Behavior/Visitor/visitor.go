package visitor

import "strings"

// 访问者接口
type IVisitor interface {
	Visit(IElement) string
}

// 成员接口
type IElement interface {
	Accept(visitor IVisitor) string
	GetValue() string
}

type ElementA struct {
	Value string
}

func (e *ElementA) Accept(visitor IVisitor) string {
	return visitor.Visit(e)
}

func (e *ElementA) GetValue() string {
	return e.Value
}

type ElementB struct {
	Value string
}

func (e *ElementB) Accept(visitor IVisitor) string {
	return visitor.Visit(e)
}

func (e *ElementB) GetValue() string {
	return e.Value
}

type Object struct {
	ElementA IElement
	ElementB IElement
}

func (o *Object) Accept(visitor IVisitor) string {
	var result []string
	result = append(result, visitor.Visit(o.ElementA), visitor.Visit(o.ElementB))
	return strings.Join(result, ",")
}

type GetValueVisitor struct{}

func (v *GetValueVisitor) Visit(element IElement) string {
	return element.GetValue()
}
