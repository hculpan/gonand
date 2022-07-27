package inputs

import (
	"github.com/hculpan/gonand/common"
)

type BasicInput struct {
	wire common.Output
}

func NewBasicInput(wire common.Output) *BasicInput {
	return &BasicInput{
		wire: wire,
	}
}

func (b *BasicInput) SetInput(value bool) {
	b.wire.SetInput(value)
}
