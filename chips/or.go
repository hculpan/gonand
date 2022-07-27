package chips

import (
	"fmt"
)

type Or struct {
	inputs map[string]InputValue
	output bool
}

func NewOr() *Or {
	result := Or{
		inputs: make(map[string]InputValue),
		output: false,
	}

	result.Reset()

	return &result
}

func (o *Or) Reset() {
	o.inputs["a"] = InputValue{Value: false, Initialized: false}
	o.inputs["b"] = InputValue{Value: false, Initialized: false}
}

func (o *Or) Ready() bool {
	for i := range o.inputs {
		if !o.inputs[i].Initialized {
			return false
		}
	}

	return true
}

func (o *Or) Result() bool {
	return o.output
}

func (o *Or) Eval() bool {
	o.output = o.inputs["a"].Value || o.inputs["b"].Value
	o.Reset()
	return o.output
}

func (o *Or) SetInput(input string, value bool) error {
	if _, ok := o.inputs[input]; ok {
		o.inputs[input] = InputValue{Value: value, Initialized: true}
	} else {
		return fmt.Errorf("invalid input '%s'", input)
	}

	if o.Ready() {
		o.Eval()
		o.Reset()
	}

	return nil
}
