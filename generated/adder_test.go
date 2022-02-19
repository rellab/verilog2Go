package generated

import (
	"fmt"
	"testing"

	"github.com/verilog2Go/src/variable"
)

func TestAdder(t *testing.T) {
	var a, b, q variable.BitArray
	a.InitBitArray(4)
	b.InitBitArray(4)
	q.InitBitArray(4)
	a.SetId("a")
	b.SetId("b")
	q.SetId("q")
	adder := NewAdder(&Adder{&a, &b, &q})
	variable.Trace("adder", []*variable.BitArray{&a, &b, &q})
	// Reset Time
	var time_counter int
	for time_counter < 100 {
		adder.Exec()
		variable.Dump(time_counter)
		time_counter++
	}

	for time_counter < 500 {
		if (time_counter % 5) == 0 {
			a.Set(a.ToInt() + 1)
		}
		if (time_counter % 50) == 0 {
			b.Set(b.ToInt() + 1)
		}

		// Evaluate DUT
		adder.Exec()
		variable.Dump(time_counter)
		time_counter++
	}
	variable.Close()
}

func TestAdderGoroutine(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	q := make(chan int)

	NewGoroutineAdder([]chan int{a, b}, []chan int{q})

	go func() {
		defer close(a)
		defer close(b)
		var time_counter int
		for time_counter < 100 {
			time_counter++
		}
		var tmpa, tmpb int
		for time_counter < 500 {
			if (time_counter % 5) == 0 {
				tmpa++
				a <- tmpa
			}
			if (time_counter % 50) == 0 {
				tmpb++
				b <- tmpb
			}
			time_counter++
		}
	}()

	for {
		select {
		case q, ok := <-q:
			if ok {
				fmt.Println(q)
			} else {
				return
			}
		}
	}
}
