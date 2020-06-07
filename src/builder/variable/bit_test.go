package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	bit := Bit{false}
	assert.Equal(t, false, bit.value)
}

func TestInitBitArray(t *testing.T) {
	var ba BitArray
	ba.InitBitArray(3, 3)
	// ba := InitBitArray(3, 3)
	assert.Equal(t, true, ba.bits[0].value)
	assert.Equal(t, true, ba.bits[1].value)
	assert.Equal(t, false, ba.bits[2].value)
}

func TestToInt(t *testing.T) {
	var ba BitArray
	ba.InitBitArray(5, 3)
	assert.Equal(t, 5, ba.ToInt())
	//オーバーフロー
	var b BitArray
	b.InitBitArray(10, 3)
	assert.Equal(t, 2, b.ToInt())
}

func TestAdd(t *testing.T) {
	var a BitArray
	a.InitBitArray(2, 3)
	var b BitArray
	b.InitBitArray(3, 3)
	assert.Equal(t, 5, a.Add(b).ToInt())
}

func TestASub(t *testing.T) {
	var a BitArray
	a.InitBitArray(3, 3)
	var b BitArray
	b.InitBitArray(2, 3)
	assert.Equal(t, 1, a.Sub(b).ToInt())
	//オーバーフロー
	assert.Equal(t, 7, b.Sub(a).ToInt())
}

func TestMul(t *testing.T) {
	var a BitArray
	a.InitBitArray(2, 3)
	var b BitArray
	b.InitBitArray(3, 3)
	assert.Equal(t, 6, a.Mul(b).ToInt())
}

func TestAssign(t *testing.T) {
	var a BitArray
	a.InitBitArray(2, 3)
	var b BitArray
	b.InitBitArray(0, 3)
	b.Assign(a)
	assert.Equal(t, 2, b.ToInt())
}
