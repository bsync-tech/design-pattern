package iterator

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type Container interface {
	GetIterator() Iterator
}

type ConcretContainerIterator struct {
	index int
	Objs  []interface{}
}

func (c *ConcretContainerIterator) HasNext() bool {
	if c.index < len(c.Objs)-1 {
		return true
	}
	return false
}

func (c *ConcretContainerIterator) Next() interface{} {
	if c.HasNext() {
		c.index += 1

		return c.Objs[c.index]
	}
	return nil
}

type ConcretContainer struct {
	Objs []interface{}
}

func (c *ConcretContainer) GetIterator() Iterator {
	return &ConcretContainerIterator{Objs: c.Objs, index: -1}
}
