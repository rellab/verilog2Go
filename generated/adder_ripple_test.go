package generated

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/verilog2Go/src/variable"
// )

// func TestAdderRipple(t *testing.T) {
// 	var a, b, q, cout variable.BitArray
// 	a.InitBitArray(4)
// 	b.InitBitArray(4)
// 	q.InitBitArray(4)
// 	cout.InitBitArray(4)
// 	adderRipple := Adder_ripple(&a, &b, &q, &cout)
// 	a.Set(0)
// 	b.Set(0)
// 	adderRipple.Exec()
// 	assert.Equal(t, 0, q.ToInt())
// 	assert.Equal(t, 0, cout.ToInt())

// 	a.Set(5)
// 	b.Set(3)
// 	adderRipple.Exec()
// 	assert.Equal(t, 8, q.ToInt())
// 	assert.Equal(t, 7, cout.ToInt())

// 	a.Set(1)
// 	b.Set(0)
// 	adderRipple.Exec()
// 	assert.Equal(t, 1, q.ToInt())
// 	assert.Equal(t, 0, cout.ToInt())

// 	a.Set(1)
// 	b.Set(1)
// 	adderRipple.Exec()
// 	assert.Equal(t, 2, q.ToInt())
// 	assert.Equal(t, 1, cout.ToInt())

// 	a.Set(0)
// 	b.Set(1)
// 	adderRipple.Exec()
// 	assert.Equal(t, 1, q.ToInt())
// 	assert.Equal(t, 0, cout.ToInt())
// }
