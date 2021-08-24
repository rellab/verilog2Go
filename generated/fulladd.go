package generated

import "github.com/verilog2Go/src/variable"

type Fulladd struct {
	A, B, CIN, Q, COUT *variable.BitArray
}

func NewFulladd(args *Fulladd) Fulladd {
	return *args
}

func NewGoroutineFulladd(in []chan int, out []chan int) *Fulladd {
	Fulladd := &Fulladd{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1)}
	go Fulladd.run(in, out)
	return Fulladd
}

func (fulladd *Fulladd) Exec() {
	fulladd.Q.Assign(fulladd.CIN.Bitxor(fulladd.B.Bitxor(*fulladd.A)))
	fulladd.COUT.Assign(fulladd.A.Bitand(*fulladd.CIN).Bitor(fulladd.CIN.Bitand(*fulladd.B).Bitor(fulladd.B.Bitand(*fulladd.A))))
}

func (fulladd *Fulladd) run(in []chan int, out []chan int) {
	defer close(out[0])
	defer close(out[1])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				fulladd.A.Set(v)
				fulladd.Exec()
				out[0] <- fulladd.Q.ToInt()
				out[1] <- fulladd.COUT.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				fulladd.B.Set(v)
				fulladd.Exec()
				out[0] <- fulladd.Q.ToInt()
				out[1] <- fulladd.COUT.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				fulladd.CIN.Set(v)
				fulladd.Exec()
				out[0] <- fulladd.Q.ToInt()
				out[1] <- fulladd.COUT.ToInt()
			} else {
				return
			}
		}
	}
}
