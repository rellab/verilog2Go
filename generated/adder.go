package generated

import "github.com/verilog2Go/src/variable"

type Adder struct {
	a, b, q *variable.BitArray
}

func NewAdder(args *Adder) Adder {
	return *args
}

func NewGoroutineAdder(in []chan variable.BitArray, out []chan variable.BitArray) *Adder {
	Adder := &Adder{variable.NewBitArray(3), variable.NewBitArray(3), variable.NewBitArray(3)}
	Adder.run(in, out)
	return Adder
}

func (Adder *Adder) Exec() {
	Adder.q.Assign(Adder.b.Add(*Adder.a))
}
