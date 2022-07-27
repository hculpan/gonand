package wires

import "github.com/hculpan/gonand/chips"

type WireDestination struct {
	name string
	chip chips.Chip
}

type Wire struct {
	outputs []WireDestination
}

func NewWireDestination(name string, chip chips.Chip) *WireDestination {
	return &WireDestination{
		name: name,
		chip: chip,
	}
}

func NewWire(outputs []WireDestination) *Wire {
	return &Wire{
		outputs: outputs,
	}
}

func NewSimpleWire(outputName string, output chips.Chip) *Wire {
	return NewWire([]WireDestination{*NewWireDestination(outputName, output)})
}

func (w *Wire) SetInput(value bool) {
	for _, c := range w.outputs {
		c.chip.SetInput(c.name, value)
	}
}
