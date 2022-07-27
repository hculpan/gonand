package inputs

import (
	"github.com/hculpan/gonand/wires"
)

type BasicInput struct {
	wire *wires.Wire
}

func NewBasicInput(wire *wires.Wire) *BasicInput {
	return &BasicInput{
		wire: wire,
	}
}

func (b *BasicInput) SetInput(value bool) {
	b.wire.SetInput(value)
}
