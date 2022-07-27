package outputs

import (
	"testing"

	"github.com/hculpan/gonand/chips"
	"github.com/hculpan/gonand/inputs"
)

func TestConsoleOutput(t *testing.T) {
	and := chips.NewAnd(NewConsoleOutput("Result:"))
	a := inputs.NewBasicInput(NewSimpleWire("a", and))
	b := inputs.NewBasicInput(NewSimpleWire("b", and))
	a.SetInput(true)
	b.SetInput(true)
	if !and.Result() {
		t.Fatalf("expected '%t' result, got '%t'", false, and.Result())
	}
}
