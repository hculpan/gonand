package chips

import (
	"fmt"

	"github.com/hculpan/gonand/common"
)

type Not struct {
	input       bool
	value       bool
	initialized bool
	output      common.Output
}

func NewNot(o common.Output) *Not {
	result := Not{
		input:  false,
		value:  false,
		output: o,
	}

	result.Reset()

	return &result
}

func (n *Not) Reset() {
	n.input = false
	n.initialized = false
}

func (n *Not) Ready() bool {
	return n.initialized
}

func (n *Not) Result() bool {
	return n.value
}

func (n *Not) Eval() bool {
	n.value = !n.input
	n.output.SetInput(n.value)
	n.Reset()
	return n.value
}

func (n *Not) SetInput(input string, value bool) error {
	if input != "in" {
		return fmt.Errorf("invalid input name '%s'", input)
	}

	n.input = value
	n.initialized = true
	n.Eval()
	n.Reset()

	return nil
}

func (n *Not) SetOutput(o common.Output) {
	n.output = o
}
