package chips

import (
	"fmt"
)

type Not struct {
	input       bool
	output      bool
	initialized bool
}

func NewNot() *Not {
	result := Not{
		input:  false,
		output: false,
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
	return n.output
}

func (n *Not) Eval() bool {
	n.output = !n.input
	n.Reset()
	return n.output
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
