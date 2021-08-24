package generated

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verilog2Go/src/variable"
)

func TestCntUnit(t *testing.T) {
	var ck, res, en, q, ca variable.BitArray
	ck.InitBitArray(1)
	res.InitBitArray(1)
	en.InitBitArray(1)
	q.InitBitArray(1)
	ca.InitBitArray(1)
	cu := NewCnt_unit(&Cnt_unit{&ck, &res, &en, &q, &ca})
	cu.Exec()
	// fmt.Println(ck.Pos)
	ck.Set(0)
	ck.Set(1)
	ck.Set(0)
	ck.Set(1)
	cu.EN.Set(1)
	ck.Set(0)
	assert.Equal(t, 0, q.ToInt())
	//クロックの立ち上がり
	ck.Set(1)
	assert.Equal(t, 1, q.ToInt())
	ck.Set(0)
	assert.Equal(t, 1, q.ToInt())
	assert.Equal(t, 1, ca.ToInt())
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

func TestCntUnitGoroutine(t *testing.T) {
	ck := make(chan int)
	res := make(chan int)
	en := make(chan int)
	q := make(chan int)
	ca := make(chan int)

	NewGoroutineCnt_unit([]chan int{ck, res, en}, []chan int{q, ca})
	go func() {
		defer close(ck)
		defer close(res)
		defer close(en)
		ck <- 0
		ck <- 1
		ck <- 0
		en <- 1
		ck <- 1
		ck <- 0
		ck <- 1
		ck <- 0
		ck <- 1
		res <- 1
	}()

	for {
		select {
		case q, ok := <-q:
			if ok {
				fmt.Printf("q: %d ", q)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		case ca, ok := <-ca:
			if ok {
				fmt.Printf("ca: %d\n", ca)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
