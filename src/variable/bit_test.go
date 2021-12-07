package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	bit := Bit{false, false, false}
	assert.Equal(t, false, bit.value)
}

func TestInitBitArray(t *testing.T) {
	var ba BitArray
	ba.InitBitArray(3)
	assert.Equal(t, false, ba.bits[0].value)
	assert.Equal(t, false, ba.bits[1].value)
	assert.Equal(t, false, ba.bits[2].value)
}

func TestGet(t *testing.T) {
	var ba BitArray
	ba.InitBitArray(3)
	ba.Set(3)
	assert.Equal(t, true, ba.Get(0).bits[0].value)
	assert.Equal(t, true, ba.Get(1).bits[0].value)
	assert.Equal(t, false, ba.Get(2).bits[0].value)
}

func TestToInt(t *testing.T) {
	var ba BitArray
	ba.InitBitArray(3)
	ba.Set(5)
	assert.Equal(t, 5, ba.ToInt())
	//オーバーフロー
	ba.Set(10)
	assert.Equal(t, 2, ba.ToInt())
}

func TestAdd(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.Set(2)
	var b BitArray
	b.InitBitArray(3)
	b.Set(3)
	assert.Equal(t, 5, a.Add(b).ToInt())
}

func TestASub(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.Set(3)
	var b BitArray
	b.InitBitArray(3)
	b.Set(2)
	assert.Equal(t, 1, a.Sub(b).ToInt())
	//オーバーフロー
	assert.Equal(t, 7, b.Sub(a).ToInt())
}

func TestMul(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.Set(2)
	var b BitArray
	b.InitBitArray(3)
	b.Set(2)
	assert.Equal(t, 4, a.Mul(b).ToInt())
}

func TestNot(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.Set(2)
	assert.Equal(t, 5, a.Not().ToInt())
	a.Set(1)
	assert.Equal(t, 6, a.Not().ToInt())
}

func TestEqual(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.Set(2)
	var b BitArray
	b.InitBitArray(3)
	b.Set(3)
	assert.Equal(t, false, a.Equal(b))
	b.Set(2)
	assert.Equal(t, true, a.Equal(b))
}

func TestAssign(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.Set(2)
	var b BitArray
	b.InitBitArray(3)
	b.Assign(a)
	assert.Equal(t, 2, b.ToInt())
}

func TestSetBits(t *testing.T) {
	var a BitArray
	a.InitBitArray(3)
	a.SetBits("3'b110")
	var b BitArray
	b.InitBitArray(3)
	b.SetBits("3'b11x")
	// fmt.Println(a)
	// fmt.Println(searchIndef(a))
	// fmt.Println(a)
	// fmt.Println(b)
	assert.Equal(t, 6, a.ToInt())
	assert.Equal(t, 0, b.ToInt())
}
