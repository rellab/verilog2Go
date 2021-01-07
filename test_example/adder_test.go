package generated

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verilog2Go/src/variable"
)

func TestAdder(t *testing.T) {
	var a, b, q variable.BitArray
	a.InitBitArray(3)
	b.InitBitArray(3)
	q.InitBitArray(3)
	adder := Adder(&a, &b, &q)
	a.Set(2)
	b.Set(3)
	adder.Exec()
	assert.Equal(t, 5, q.ToInt())
}
