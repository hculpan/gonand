package outputs

type NullOutput struct {
}

func NewNullOutput() *NullOutput {
	return &NullOutput{}
}

func (n *NullOutput) SetInput(value bool) {
	// do nothing
}
