package facade

import "fmt"

type Cpu struct{}

func (c *Cpu) LoadInstruction() string {
	return "cpu load instructions"
}

func (c *Cpu) Jump() string {
	return "cpu jump"
}

func (c *Cpu) Execute() string {
	return "cpu execute"
}

type Memory struct{}

func (m *Memory) LoadData() string {
	return "memory load data"
}

type Disk struct{}

func (d *Disk) Read() string {
	return "disk read"
}

type ComputerFacade struct {
	Cpu    *Cpu
	Memory *Memory
	Disk   *Disk
}

func (c *ComputerFacade) Run() string {
	return fmt.Sprintf("%s,%s,%s,%s", c.Cpu.LoadInstruction(), c.Memory.LoadData(), c.Disk.Read(), c.Cpu.Execute())
}
