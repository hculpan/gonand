package inputs

import (
	"fmt"
	"testing"

	"github.com/hculpan/gonand/chips"
	"github.com/hculpan/gonand/outputs"
)

func TestAnd(t *testing.T) {
	if ok := testAndWithResult(true, true, true); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testAndWithResult(true, false, false); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testAndWithResult(false, true, false); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testAndWithResult(false, false, false); ok != nil {
		t.Fatalf(ok.Error())
	}
}

func testAndWithResult(aValue bool, bValue bool, expected bool) error {
	and := chips.NewAnd(outputs.NewNullOutput())
	a := NewBasicInput(outputs.NewSimpleWire("a", and))
	b := NewBasicInput(outputs.NewSimpleWire("b", and))
	a.SetInput(aValue)
	b.SetInput(bValue)
	if and.Result() != expected {
		return fmt.Errorf("expected '%t' result, got '%t'", expected, and.Result())
	}

	return nil
}
