package generated

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verilog2Go/src/variable"
)

func TestSaikoro(t *testing.T) {
	var ck, reset, enable, lamp, cnt variable.BitArray
	ck.InitBitArray(1)
	reset.InitBitArray(1)
	enable.InitBitArray(1)
	lamp.InitBitArray(7)
	cnt.InitBitArray(3)
	NewSaikoro(&Saikoro{&ck, &reset, &enable, &lamp, &cnt})
	// sa.Exec()
	reset.Set(0)
	enable.Set(1)
	ck.Set(0)
	ck.Set(1)
	assert.Equal(t, 8, lamp.ToInt())
	ck.Set(0)
	ck.Set(1)
	assert.Equal(t, 65, lamp.ToInt())
	ck.Set(0)
	ck.Set(1)
	assert.Equal(t, 28, lamp.ToInt())
	ck.Set(0)
	ck.Set(1)
	assert.Equal(t, 85, lamp.ToInt())
	ck.Set(0)
	ck.Set(1)
	assert.Equal(t, 93, lamp.ToInt())
	reset.Set(1)
	// reset.Pos.Always()
	assert.Equal(t, 8, lamp.ToInt())
}
