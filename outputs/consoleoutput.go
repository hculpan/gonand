package outputs

import "fmt"

type ConsoleOutput struct {
	message string
}

func NewConsoleOutput(message string) *ConsoleOutput {
	return &ConsoleOutput{message: message}
}

func (c *ConsoleOutput) SetInput(value bool) {
	fmt.Printf("%s %t\n", c.message, value)
}
