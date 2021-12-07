package generated

import "github.com/verilog2Go/src/variable"

type Alu struct {
	a, b, f, s, x, y *variable.BitArray
}

func NewAlu(args *Alu) Alu {
	return *args
}

func NewGoroutineAlu(in []chan int, out []chan int) *Alu {
	alu := &Alu{variable.NewBitArray(16), variable.NewBitArray(16), variable.NewBitArray(5), variable.NewBitArray(16), variable.NewBitArray(16), variable.NewBitArray(16)}
	go alu.run(in, out)
	return alu
}

func (alu *Alu) Exec() {
	alu.x.Assign(variable.CreateBitArray(16, 32768).Add(*alu.a))
	alu.y.Assign(variable.CreateBitArray(16, 32768).Add(*alu.b))
}

func (alu *Alu) run(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				alu.a.Set(v)
				bitArrays1 := alu.PreAlways1()
				alu.Always1(bitArrays1)
				alu.Exec()
				out[0] <- alu.s.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				alu.b.Set(v)
				bitArrays1 := alu.PreAlways1()
				alu.Always1(bitArrays1)
				alu.Exec()
				out[0] <- alu.s.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				alu.f.Set(v)
				bitArrays1 := alu.PreAlways1()
				alu.Always1(bitArrays1)
				alu.Exec()
				out[0] <- alu.s.ToInt()
			} else {
				return
			}
		}
	}
}

func (Alu *Alu) PreAlways1() []variable.BitArray {
	var1 := *variable.CreateBitArray(16, Alu.a.ToInt())
	var2 := *variable.CreateBitArray(16, Alu.b.ToInt())
	var3 := *variable.CreateBitArray(5, Alu.f.ToInt())
	return []variable.BitArray{var1, var2, var3}
}
func (Alu *Alu) Always1(vars []variable.BitArray) {
}
