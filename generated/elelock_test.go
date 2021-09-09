package generated

import (
	"fmt"
	"testing"

	"github.com/verilog2Go/src/variable"
)

func TestElelock(t *testing.T) {
	var clk, reset, close, tenkey, lock, ke1, ke2, match, key_enbl, SECRET_3, SECRET_2, SECRET_1, SECRET_0, key0, key1, key2, key3 variable.BitArray
	key := make([]*variable.BitArray, 4)
	clk.InitBitArray(1)
	reset.InitBitArray(1)
	close.InitBitArray(1)
	tenkey.InitBitArray(10)
	lock.InitBitArray(1)
	ke1.InitBitArray(1)
	ke2.InitBitArray(1)
	match.InitBitArray(1)
	key_enbl.InitBitArray(1)
	SECRET_3.InitBitArray(4)
	SECRET_2.InitBitArray(4)
	SECRET_1.InitBitArray(4)
	SECRET_0.InitBitArray(4)
	key3.InitBitArray(4)
	key2.InitBitArray(4)
	key1.InitBitArray(4)
	key0.InitBitArray(4)
	key[0] = &key0
	key[1] = &key1
	key[2] = &key2
	key[3] = &key3
	elelock := NewElelock(&Elelock{&clk, &reset, &close, &tenkey, &lock, &ke1, &ke2, &match, &key_enbl, &SECRET_3, &SECRET_2, &SECRET_1, &SECRET_0, key})
	// Reset Time
	var time_counter int
	for time_counter < 100 {
		elelock.Exec()
		time_counter++
	}

	for time_counter < 500 {
		if (time_counter % 5) == 0 {
			clk.Set(clk.ToInt() + 1)
		}
		if (time_counter % 50) == 0 {
			reset.Set(reset.ToInt() + 1)
		}

		// Evaluate DUT
		elelock.Exec()
		fmt.Println(lock.ToInt())
		time_counter++
	}
	// a.Set(2)
	// b.Set(3)
	// adder.Exec()
	// assert.Equal(t, 5, q.ToInt())
	// a.Set(3)
	// b.Set(4)
	// adder.Exec()
	// assert.Equal(t, 7, q.ToInt())
	// a.Set(5)
	// b.Set(5)
	// adder.Exec()
	// assert.Equal(t, 2, q.ToInt())
}

func TestElelockGoroutine(t *testing.T) {
	clk := make(chan int)
	reset := make(chan int)
	clos := make(chan int)
	tenkey := make(chan int)
	lock := make(chan int)

	NewGoroutineAdder([]chan int{clk, reset, clos, tenkey}, []chan int{lock})

	go func() {
		defer close(clk)
		defer close(reset)
		defer close(clos)
		defer close(tenkey)
		var time_counter int
		for time_counter < 100 {
			time_counter++
		}
		var tmpa, tmpb int
		for time_counter < 500 {
			if (time_counter % 5) == 0 {
				tmpa++
				clk <- tmpa
			}
			if (time_counter % 50) == 0 {
				tmpb++
				reset <- tmpb
			}
			time_counter++
		}
		// a <- 1
		// b <- 2
		// b <- 4
		// a <- 7
		// b <- 11
	}()

	for {
		select {
		case q, ok := <-lock:
			if ok {
				fmt.Println(q)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
