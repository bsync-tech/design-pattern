package command

type RequestInterface interface {
	Get() string
}

type GetRequest struct {
	Method string
}

func (g *GetRequest) Get() string {
	return g.Method
}

type PostRequest struct {
	Method string
}

func (p *PostRequest) Get() string {
	return p.Method
}

type Request struct {
	Request RequestInterface
}

func Action(r *Request) string {
	if r.Request.Get() == "GET" {
		return "get method executed"
	} else if r.Request.Get() == "POST" {
		return "post method executed"
	}
	return "invalid method"
}
