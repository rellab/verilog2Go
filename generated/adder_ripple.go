package generated

import "github.com/verilog2Go/src/variable"

type adder_ripple struct {
	a, b, q, cout variable.BitArray
}

func Adder_ripple(args *adder_ripple) adder_ripple {
	return *args
}

func (adder_ripple *adder_ripple) Exec() {
	// add0 := Fulladd(&fulladd{Q: adder_ripple.q.Get(0), COUT: adder_ripple.cout.Get(0), A: adder_ripple.a.Get(0), B: adder_ripple.b.Get(0), CIN: variable.CreateBitArray(1, 0)})
	// add0.Exec()
	// add1 := Fulladd(&fulladd{COUT: adder_ripple.cout.Get(1), A: adder_ripple.a.Get(1), B: adder_ripple.b.Get(1), CIN: adder_ripple.cout.Get(0), Q: adder_ripple.q.Get(1)})
	// add1.Exec()
	// add2 := Fulladd(&fulladd{Q: adder_ripple.q.Get(2), COUT: adder_ripple.cout.Get(2), A: adder_ripple.a.Get(2), B: adder_ripple.b.Get(2), CIN: adder_ripple.cout.Get(1)})
	// add2.Exec()
	// add3 := Fulladd(&fulladd{COUT: adder_ripple.cout.Get(3), A: adder_ripple.a.Get(3), B: adder_ripple.b.Get(3), CIN: adder_ripple.cout.Get(2), Q: adder_ripple.q.Get(3)})
	// add3.Exec()
}
