package chips

import (
	"fmt"

	"github.com/hculpan/gonand/common"
)

type And struct {
	inputs map[string]common.InputValue
	value  bool
	output common.Output
}

func NewAnd(output common.Output) *And {
	result := And{
		inputs: make(map[string]common.InputValue),
		value:  false,
		output: output,
	}

	result.Reset()

	return &result
}

func (a *And) Reset() {
	a.inputs["a"] = common.InputValue{Value: false, Initialized: false}
	a.inputs["b"] = common.InputValue{Value: false, Initialized: false}
}

func (a *And) Ready() bool {
	for i := range a.inputs {
		if !a.inputs[i].Initialized {
			return false
		}
	}

	return true
}

func (a *And) Result() bool {
	return a.value
}

func (a *And) Eval() bool {
	a.value = a.inputs["a"].Value && a.inputs["b"].Value
	a.output.SetInput(a.value)
	a.Reset()
	return a.value
}

func (a *And) SetInput(input string, value bool) error {
	if _, ok := a.inputs[input]; ok {
		a.inputs[input] = common.InputValue{Value: value, Initialized: true}
	} else {
		return fmt.Errorf("invalid input '%s'", input)
	}

	if a.Ready() {
		a.Eval()
		a.Reset()
	}

	return nil
}

func (a *And) SetOutput(o common.Output) {
	a.output = o
}
