package generated

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/verilog2Go/src/variable"
// )

// func TestCounter(t *testing.T) {
// 	var ck, res, q variable.BitArray
// 	ck.InitBitArray(1)
// 	res.InitBitArray(1)
// 	q.InitBitArray(4)
// 	cu := NewCounter(&Counter{&ck, &res, &q})
// 	cu.Exec()
// 	// fmt.Println(ck.Pos)
// 	ck.Set(0)
// 	ck.Set(1)
// 	ck.Set(0)
// 	ck.Set(1)
// 	ck.Set(0)
// }

// func TestCounterGoroutine(t *testing.T) {
// 	ck := make(chan int)
// 	res := make(chan int)
// 	q := make(chan int)
// 	NewGoroutineCounter([]chan int{ck, res}, []chan int{q})
// 	go func() {
// 		defer close(ck)
// 		defer close(res)
// 		ck <- 0
// 		ck <- 1
// 		ck <- 0
// 		ck <- 1
// 		ck <- 0
// 		// ck <- 1
// 		// ck <- 0
// 		// ck <- 1
// 		// res <- 1
// 	}()

// 	for {
// 		select {
// 		case q, ok := <-q:
// 			if ok {
// 				fmt.Printf("q: %d ", q)
// 			} else {
// 				// アウトプットチャンネルがクローズされたら終了
// 				return
// 			}
// 		}
// 	}
// }
