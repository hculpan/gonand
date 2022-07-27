package inputs

import (
	"fmt"
	"testing"

	"github.com/hculpan/gonand/chips"
	"github.com/hculpan/gonand/outputs"
)

func TestNand(t *testing.T) {
	if ok := testNandWithResult(true, true, false); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testNandWithResult(true, false, true); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testNandWithResult(false, true, true); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testNandWithResult(false, false, true); ok != nil {
		t.Fatalf(ok.Error())
	}
}

func testNandWithResult(aValue bool, bValue bool, expected bool) error {
	nand := chips.NewNand(outputs.NewNullOutput())
	a := NewBasicInput(outputs.NewSimpleWire("a", nand))
	b := NewBasicInput(outputs.NewSimpleWire("b", nand))
	a.SetInput(aValue)
	b.SetInput(bValue)
	if nand.Result() != expected {
		return fmt.Errorf("expected '%t' result, got '%t'", expected, nand.Result())
	}

	return nil
}
