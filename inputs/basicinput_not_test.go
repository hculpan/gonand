package inputs

import (
	"fmt"
	"testing"

	"github.com/hculpan/gonand/chips"
	"github.com/hculpan/gonand/wires"
)

func TestNot(t *testing.T) {
	if ok := testNotWithResult(true, false); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testNotWithResult(false, true); ok != nil {
		t.Fatalf(ok.Error())
	}
}

func testNotWithResult(inValue bool, expected bool) error {
	not := chips.NewNot()
	in := NewBasicInput(wires.NewSimpleWire("in", not))
	in.SetInput(inValue)
	if not.Result() != expected {
		return fmt.Errorf("expected '%t' result, got '%t'", expected, not.Result())
	}

	return nil
}
