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
