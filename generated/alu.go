package generated

import "github.com/verilog2Go/src/variable"

type Alu struct {
	a, b, f, s, x, y *variable.BitArray
}

func NewAlu() Alu {
	args := &Alu{variable.NewBitArray(16), variable.NewBitArray(16), variable.NewBitArray(5), variable.NewBitArray(16), variable.NewBitArray(16), variable.NewBitArray(16)}
	return *args
}

func NewGoroutineAlu(in []chan int, out []chan int) *Alu {
	alu := &Alu{variable.NewBitArray(16), variable.NewBitArray(16), variable.NewBitArray(5), variable.NewBitArray(16), variable.NewBitArray(16), variable.NewBitArray(16)}
	go alu.start(in, out)
	return alu
}

func (alu *Alu) Exec() {
	alu.x.Assign(variable.CreateBits("16'h8000").Add(*alu.a))
	alu.y.Assign(variable.CreateBits("16'h8000").Add(*alu.b))
}

func (alu *Alu) start(in []chan int, out []chan int) {
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
	var4 := *variable.CreateBitArray(16, Alu.x.ToInt())
	var5 := *variable.CreateBitArray(16, Alu.y.ToInt())
	return []variable.BitArray{var1, var2, var3, var4, var5}
}

func (Alu *Alu) Always1(vars []variable.BitArray) {
	switch Alu.f.ToInt() {
	case 0:
		Alu.s.Assign(Alu.a.Add(*Alu.b))

	case 1:
		Alu.s.Assign(Alu.a.Sub(*Alu.b))

	case 2:
		Alu.s.Assign(Alu.a.Mul(*Alu.b))

	case 3:
		Alu.s.Assign(Alu.a.SHL(*Alu.b))

	case 4:
		Alu.s.Assign(Alu.a.SHR(*Alu.b))

	case 5:
		Alu.s.Assign(Alu.a.Bitand(*Alu.b))

	case 6:
		Alu.s.Assign(Alu.a.Bitor(*Alu.b))

	case 7:
		Alu.s.Assign(Alu.a.Bitxor(*Alu.b))

	case 8:
		Alu.s.Assign(Alu.a.And(*Alu.b))

	case 9:
		Alu.s.Assign(Alu.a.Or(*Alu.b))

	case 10:
		Alu.s.Assign(Alu.a.Equal(*Alu.b))

	case 11:
		Alu.s.Assign(Alu.a.NE(*Alu.b))

	case 12:
		Alu.s.Assign(Alu.x.GE(*Alu.y))

	case 13:
		Alu.s.Assign(Alu.x.LE(*Alu.y))

	case 14:
		Alu.s.Assign(Alu.x.GT(*Alu.y))

	case 15:
		Alu.s.Assign(Alu.x.LT(*Alu.y))

	case 16:
		Alu.s.Assign(Alu.a.Neg())

	case 17:
		Alu.s.Assign(Alu.a.Bnot())

	case 18:
		Alu.s.Assign(Alu.a.Not())

	default:
		Alu.s.Assign(*variable.CreateBits("16'hxxxx"))

	}
}
