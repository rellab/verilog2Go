package generated

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/verilog2Go/src/variable"
// )

// var ck, res, en, q, ca variable.BitArray
// var cu cnt_unit

// func TestMain(m *testing.M) {
// 	//テストに必要な変数の初期化
// 	ck.InitBitArray(1)
// 	res.InitBitArray(1)
// 	en.InitBitArray(1)
// 	q.InitBitArray(1)
// 	ca.InitBitArray(1)
// 	cu = Cnt_unit(&ck, &res, &en, &q, &ca)

// 	//テストケース実行
// 	m.Run()
// }

// func TestCntUnit(t *testing.T) {
// 	cu.Exec()
// 	// fmt.Println(ck.Pos)
// 	ck.Set(0)
// 	ck.Set(1)
// 	ck.Set(0)
// 	ck.Set(1)
// 	cu.EN.Set(1)
// 	ck.Set(0)
// 	assert.Equal(t, 0, q.ToInt())
// 	//クロックの立ち上がり
// 	ck.Set(1)
// 	assert.Equal(t, 1, q.ToInt())
// 	ck.Set(0)
// 	assert.Equal(t, 1, q.ToInt())
// 	assert.Equal(t, 1, ca.ToInt())
// 	//クロックの立ち上がり
// 	ck.Set(1)
// 	assert.Equal(t, 0, q.ToInt())
// 	assert.Equal(t, 0, ca.ToInt())
// 	ck.Set(0)
// 	//クロックの立ち上がり
// 	ck.Set(1)
// 	assert.Equal(t, 1, q.ToInt())
// 	assert.Equal(t, 1, ca.ToInt())
// 	//リセットの立ち上がり
// 	res.Set(1)
// 	assert.Equal(t, 0, q.ToInt())
// 	assert.Equal(t, 0, ca.ToInt())
// }
