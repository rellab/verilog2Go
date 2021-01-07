package generated

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verilog2Go/src/variable"
)

var ck, res, q, ca variable.BitArray
var cu counter_unit

func TestMain(m *testing.M) {
	//テストに必要な変数の初期化
	ck.InitBitArray(1)
	res.InitBitArray(1)
	q.InitBitArray(4)
	ca.InitBitArray(4)
	cu = Counter_unit(&ck, &res, &q, &ca)

	//テストケース実行
	m.Run()
}

func TestCounterUnit(t *testing.T) {
	cu.Exec()
	// fmt.Println(cu.ck.Pos)
	cu.ck.Set(0)
	// fmt.Printf("q: %d, ca: %d\n", cu.q.ToInt(), cu.ca.ToInt())
	cu.ck.Set(1)
	// fmt.Printf("q: %d, ca: %d\n", cu.q.ToInt(), cu.ca.ToInt())
	cu.ck.Set(0)
	assert.Equal(t, 0, q.ToInt())
	// クロックの立ち上がり
	cu.ck.Set(1)
	// fmt.Printf("q: %d, ca: %d\n", cu.q.ToInt(), cu.ca.ToInt())
	assert.Equal(t, 1, q.ToInt())
	// fmt.Printf("q: %d, ca: %d\n", cu.q.ToInt(), cu.ca.ToInt())
	assert.Equal(t, 1, ca.ToInt())
	ck.Set(0)
	//クロックの立ち上がり
	ck.Set(1)
	assert.Equal(t, 0, q.ToInt())
	assert.Equal(t, 0, ca.ToInt())
	ck.Set(0)
	//クロックの立ち上がり
	ck.Set(1)
	assert.Equal(t, 1, q.ToInt())
	assert.Equal(t, 1, ca.ToInt())
	//リセットの立ち上がり
	res.Set(1)
	assert.Equal(t, 0, q.ToInt())
	assert.Equal(t, 0, ca.ToInt())
}
