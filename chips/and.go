package chips

import (
	"fmt"
)

type And struct {
	inputs map[string]InputValue
	output bool
}

func NewAnd() *And {
	result := And{
		inputs: make(map[string]InputValue),
		output: false,
	}

	result.Reset()

	return &result
}

func (a *And) Reset() {
	a.inputs["a"] = InputValue{Value: false, Initialized: false}
	a.inputs["b"] = InputValue{Value: false, Initialized: false}
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
	return a.output
}

func (a *And) Eval() bool {
	a.output = a.inputs["a"].Value && a.inputs["b"].Value
	a.Reset()
	return a.output
}

func (a *And) SetInput(input string, value bool) error {
	if _, ok := a.inputs[input]; ok {
		a.inputs[input] = InputValue{Value: value, Initialized: true}
	} else {
		return fmt.Errorf("invalid input '%s'", input)
	}

	if a.Ready() {
		a.Eval()
		a.Reset()
	}

	return nil
}
