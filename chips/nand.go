package chips

import (
	"fmt"

	"github.com/hculpan/gonand/common"
)

type Nand struct {
	inputs map[string]common.InputValue
	value  bool
	output common.Output
}

func NewNand(output common.Output) *Nand {
	result := Nand{
		inputs: make(map[string]common.InputValue),
		value:  false,
		output: output,
	}

	result.Reset()

	return &result
}

func (n *Nand) Reset() {
	n.inputs["a"] = common.InputValue{Value: false, Initialized: false}
	n.inputs["b"] = common.InputValue{Value: false, Initialized: false}
}

func (n *Nand) Ready() bool {
	for i := range n.inputs {
		if !n.inputs[i].Initialized {
			return false
		}
	}

	return true
}

func (n *Nand) Result() bool {
	return n.value
}

func (n *Nand) Eval() bool {
	n.value = !(n.inputs["a"].Value && n.inputs["b"].Value)
	n.output.SetInput(n.value)
	n.Reset()
	return n.value
}

func (n *Nand) SetInput(input string, value bool) error {
	if _, ok := n.inputs[input]; ok {
		n.inputs[input] = common.InputValue{Value: value, Initialized: true}
	} else {
		return fmt.Errorf("invalid input '%s'", input)
	}

	if n.Ready() {
		n.Eval()
		n.Reset()
	}

	return nil
}

func (n *Nand) SetOutput(o common.Output) {
	n.output = o
}
