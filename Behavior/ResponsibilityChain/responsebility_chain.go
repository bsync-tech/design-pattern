package rchain

import "fmt"

type IHandle interface {
	SetNext(next IHandle)
	GetNext() IHandle
	HandleRequest(request string) string
	GetName() string
}

type HandlerA struct {
	Name string
	Next IHandle
}

func (h *HandlerA) GetName() string {
	return h.Name
}
func (h *HandlerA) SetNext(next IHandle) {
	h.Next = next
}
func (h *HandlerA) GetNext() IHandle {
	return h.Next
}
func (h *HandlerA) HandleRequest(request string) string {
	if request == "need A" {
		return fmt.Sprintf("request %s processed by %s", request, h.Name)
	} else {
		if h.GetNext() != nil {
			return h.GetNext().HandleRequest(request)
		}
	}

	return "no handler after A"
}

type HandlerB struct {
	Name string
	Next IHandle
}

func (h *HandlerB) GetName() string {
	return h.Name
}
func (h *HandlerB) SetNext(next IHandle) {
	h.Next = next
}
func (h *HandlerB) GetNext() IHandle {
	return h.Next
}
func (h *HandlerB) HandleRequest(request string) string {

	if request == "need B" {
		return fmt.Sprintf("request %s processed by %s", request, h.Name)
	} else {
		if h.GetNext() != nil {
			h.GetNext().HandleRequest(request)
		}
	}
	return "no handler after B"
}
