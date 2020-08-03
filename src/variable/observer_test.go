package variable_test

import (
	"fmt"
	"testing"

	"github.com/verilog2Go/src/builder/variable"
)

type TestModule struct {
	a variable.BitArray
}

func (tm TestModule) Always() {
	fmt.Println("Posedge")
}

func TestPosedge(t *testing.T) {
	var ba variable.BitArray
	ba.InitBitArray(3)
	ba.Set(3)
	var tm TestModule
	tm.a = ba
	tm.a.AddPosedgeObserver(tm)

	var b variable.BitArray
	b.InitBitArray(3)
	b.Set(2)
	tm.a.Add(b)
}
