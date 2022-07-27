package chips

import (
	"fmt"

	"github.com/hculpan/gonand/common"
)

type Or struct {
	inputs map[string]common.InputValue
	value  bool
	output common.Output
}

func NewOr(o common.Output) *Or {
	result := Or{
		inputs: make(map[string]common.InputValue),
		value:  false,
		output: o,
	}

	result.Reset()

	return &result
}

func (o *Or) Reset() {
	o.inputs["a"] = common.InputValue{Value: false, Initialized: false}
	o.inputs["b"] = common.InputValue{Value: false, Initialized: false}
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
	return o.value
}

func (o *Or) Eval() bool {
	o.value = o.inputs["a"].Value || o.inputs["b"].Value
	o.output.SetInput(o.value)
	o.Reset()
	return o.value
}

func (o *Or) SetInput(input string, value bool) error {
	if _, ok := o.inputs[input]; ok {
		o.inputs[input] = common.InputValue{Value: value, Initialized: true}
	} else {
		return fmt.Errorf("invalid input '%s'", input)
	}

	if o.Ready() {
		o.Eval()
		o.Reset()
	}

	return nil
}

func (o *Or) SetOutput(out common.Output) {
	o.output = out
}
