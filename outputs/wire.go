package outputs

import "github.com/hculpan/gonand/common"

type WireDestination struct {
	name string
	chip common.Chip
}

type Wire struct {
	outputs []WireDestination
}

func NewWireDestination(name string, chip common.Chip) *WireDestination {
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

func NewSimpleWire(outputName string, output common.Chip) *Wire {
	return NewWire([]WireDestination{*NewWireDestination(outputName, output)})
}

func (w *Wire) SetInput(value bool) {
	for _, c := range w.outputs {
		c.chip.SetInput(c.name, value)
	}
}
