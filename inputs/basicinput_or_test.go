package inputs

import (
	"fmt"
	"testing"

	"github.com/hculpan/gonand/chips"
	"github.com/hculpan/gonand/outputs"
)

func TestOr(t *testing.T) {
	if ok := testOrWithResult(true, true, true); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testOrWithResult(true, false, true); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testOrWithResult(false, true, true); ok != nil {
		t.Fatalf(ok.Error())
	}
	if ok := testOrWithResult(false, false, false); ok != nil {
		t.Fatalf(ok.Error())
	}
}

func testOrWithResult(aValue bool, bValue bool, expected bool) error {
	or := chips.NewOr(outputs.NewNullOutput())
	a := NewBasicInput(outputs.NewSimpleWire("a", or))
	b := NewBasicInput(outputs.NewSimpleWire("b", or))
	a.SetInput(aValue)
	b.SetInput(bValue)
	if or.Result() != expected {
		return fmt.Errorf("expected '%t' result, got '%t'", expected, or.Result())
	}

	return nil
}
