package chips

type InputValue struct {
	Value       bool
	Initialized bool
}

type Chip interface {
	Reset()
	Ready() bool
	SetInput(input string, value bool) error
	Result() bool
}
