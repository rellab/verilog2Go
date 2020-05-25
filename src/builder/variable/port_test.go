package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitPort(t *testing.T) {
	var port Port
	port.InitBitArray(3)
	assert.Equal(t, false, port.bits[0].value)
	assert.Equal(t, false, port.bits[1].value)
	assert.Equal(t, false, port.bits[2].value)
}

func TestAddPort(t *testing.T) {
	var a Port
	a.InitBitArray(3)
	a.Set(2)
	var b Port
	b.InitBitArray(3)
	b.Set(3)
	assert.Equal(t, 5, a.Add(b).ToInt())
}
