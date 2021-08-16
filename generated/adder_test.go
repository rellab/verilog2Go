package generated

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verilog2Go/src/variable"
)

func TestAdder(t *testing.T) {
	var a, b, q variable.BitArray
	a.InitBitArray(3)
	b.InitBitArray(3)
	q.InitBitArray(3)
	adder := NewAdder(&Adder{&a, &b, &q})
	a.Set(2)
	b.Set(3)
	adder.Exec()
	assert.Equal(t, 5, q.ToInt())
	a.Set(3)
	b.Set(4)
	adder.Exec()
	assert.Equal(t, 7, q.ToInt())
	a.Set(5)
	b.Set(5)
	adder.Exec()
	assert.Equal(t, 2, q.ToInt())
}

func TestAdderGoroutine(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	q := make(chan int)

	NewGoroutineAdder([]chan int{a, b}, []chan int{q})

	go func() {
		defer close(a)
		defer close(b)
		a <- 1
		b <- 2
		b <- 4
		a <- 7
		b <- 11
	}()

	for {
		select {
		case q, ok := <-q:
			if ok {
				fmt.Println(q)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
