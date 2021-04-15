package generated

import "github.com/verilog2Go/src/variable"

type adder struct {
	a, b, q variable.BitArray
}

func Adder(a *variable.BitArray, b *variable.BitArray, q *variable.BitArray) adder {
	p := new(adder)
	p.a = *a
	p.b = *b
	p.q = *q
	return *p
}

func (adder *adder) Exec() {
	adder.q.Assign(adder.b.Add(adder.a))
}
