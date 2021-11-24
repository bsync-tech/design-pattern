package builder

import (
	"fmt"
)

type Computer struct {
	cpu      string
	keyboard string
	mouse    string
}

func (c *Computer) setCpu(cpu string) {
	c.cpu = cpu
}
func (c *Computer) setKeyboard(keyboard string) {
	c.keyboard = keyboard
}
func (c *Computer) setMouse(mouse string) {
	c.mouse = mouse
}
func (c Computer) String() string {
	return fmt.Sprintf("%s %s %s", c.cpu, c.keyboard, c.mouse)
}

type ComputerBuilder interface {
	BuildCpu()
	BuildKeyboard()
	BuildMouse()
	GetObject() *Computer
}

type HighPerformanceComputerBuilder struct {
	computer Computer
}

func (c *HighPerformanceComputerBuilder) BuildCpu() {
	c.computer.setCpu("amd 3995wx")
}
func (c *HighPerformanceComputerBuilder) BuildKeyboard() {
	c.computer.setKeyboard("ps-2")
}
func (c *HighPerformanceComputerBuilder) BuildMouse() {
	c.computer.setMouse("logitech")
}
func (c *HighPerformanceComputerBuilder) GetObject() *Computer {
	return &c.computer
}

type Director struct {
	builder ComputerBuilder
}

func (d *Director) SetBuilder(builder ComputerBuilder) {
	d.builder = builder
}
func (d *Director) Build() *Computer {
	d.builder.BuildCpu()
	d.builder.BuildKeyboard()
	d.builder.BuildMouse()
	return d.builder.GetObject()
}
