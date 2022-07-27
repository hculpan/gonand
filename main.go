package main

import (
	"github.com/hculpan/gonand/chips"
	"github.com/hculpan/gonand/inputs"
	"github.com/hculpan/gonand/outputs"
)

func main() {
	not := chips.NewNot(outputs.NewConsoleOutput("NAND result:"))
	and := chips.NewAnd(outputs.NewSimpleWire("in", not))
	a := inputs.NewBasicInput(outputs.NewSimpleWire("a", and))
	b := inputs.NewBasicInput(outputs.NewSimpleWire("b", and))
	a.SetInput(true)
	b.SetInput(true)

	a.SetInput(true)
	b.SetInput(false)
}
