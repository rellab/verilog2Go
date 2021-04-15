package generated

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verilog2Go/src/variable"
)

// func TestFulladd(t *testing.T) {
// 	var a, b, q, cin, cout variable.BitArray
// 	a.InitBitArray(1)
// 	b.InitBitArray(1)
// 	cin.InitBitArray(1)
// 	q.InitBitArray(1)
// 	cout.InitBitArray(1)
// 	fulladd := Fulladd(&a, &b, &cin, &q, &cout)
// 	a.Set(0)
// 	b.Set(0)
// 	cin.Set(0)
// 	fulladd.Exec()
// 	assert.Equal(t, 0, q.ToInt())
// 	assert.Equal(t, 0, cout.ToInt())

// 	a.Set(1)
// 	b.Set(0)
// 	cin.Set(0)
// 	fulladd.Exec()
// 	assert.Equal(t, 1, q.ToInt())
// 	assert.Equal(t, 0, cout.ToInt())

// 	a.Set(1)
// 	b.Set(1)
// 	cin.Set(0)
// 	fulladd.Exec()
// 	assert.Equal(t, 0, q.ToInt())
// 	assert.Equal(t, 1, cout.ToInt())

// 	a.Set(1)
// 	b.Set(1)
// 	cin.Set(1)
// 	fulladd.Exec()
// 	assert.Equal(t, 1, q.ToInt())
// 	assert.Equal(t, 1, cout.ToInt())

// 	a.Set(0)
// 	b.Set(1)
// 	cin.Set(1)
// 	fulladd.Exec()
// 	assert.Equal(t, 0, q.ToInt())
// 	assert.Equal(t, 1, cout.ToInt())
// }

func TestFulladdInstance(t *testing.T) {
	var a, b, q, cin, cout variable.BitArray
	a.InitBitArray(1)
	b.InitBitArray(1)
	cin.InitBitArray(1)
	q.InitBitArray(1)
	cout.InitBitArray(1)
	fa := Fulladd(&fulladd{A: a, B: b, CIN: cin, Q: q, COUT: cout})
	a.Set(0)
	b.Set(0)
	cin.Set(0)
	fa.Exec()
	assert.Equal(t, 0, q.ToInt())
	assert.Equal(t, 0, cout.ToInt())

	a.Set(1)
	b.Set(0)
	cin.Set(0)
	fa.Exec()
	assert.Equal(t, 1, q.ToInt())
	assert.Equal(t, 0, cout.ToInt())
}
