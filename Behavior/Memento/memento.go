package memento

type Memento struct {
	snapshot string
}

func (mem *Memento) SetSnapshot(state string) {
	mem.snapshot = state
}
func (mem *Memento) GetSnapshot() string {
	return mem.snapshot
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}
func (o *Originator) GetState() string {
	return o.state
}
func (o *Originator) CreateMemento() Memento {
	return Memento{snapshot: o.state}
}
func (o *Originator) RestoreMemento(mem Memento) {
	o.SetState(mem.GetSnapshot())
}

type MemoManager struct {
	memento *Memento
}

func (c *MemoManager) SetMemento(mem *Memento) {
	c.memento = mem
}
func (c *MemoManager) GetMemento() *Memento {
	return c.memento
}
