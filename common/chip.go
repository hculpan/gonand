package common

type InputValue struct {
	Value       bool
	Initialized bool
}

type Chip interface {
	Reset()
	Ready() bool
	SetInput(input string, value bool) error
	Eval() bool
	Result() bool
	SetOutput(o Output)
}
