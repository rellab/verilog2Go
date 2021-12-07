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
	alu.x.Assign(variable.CreateBits("16'h8000").Add(*alu.a))
	alu.y.Assign(variable.CreateBits("16'h8000").Add(*alu.b))
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
	var4 := *variable.CreateBitArray(16, Alu.x.ToInt())
	var5 := *variable.CreateBitArray(16, Alu.y.ToInt())
	return []variable.BitArray{var1, var2, var3, var4, var5}
}
func (Alu *Alu) Always1(vars []variable.BitArray) {
	switch Alu.f.ToInt() {
	case 0:
		Alu.s.Substitute(Alu.a.Add(*Alu.b))
		break
	case 1:
		Alu.s.Substitute(Alu.a.Sub(*Alu.b))
		break
	case 2:
		Alu.s.Substitute(Alu.a.Mul(*Alu.b))
		break
	case 3:
		Alu.s.Substitute(Alu.a.SHL(*Alu.b))
		break
	case 4:
		Alu.s.Substitute(Alu.a.SHR(*Alu.b))
		break
	case 5:
		Alu.s.Substitute(Alu.a.Bitand(*Alu.b))
		break
	case 6:
		Alu.s.Substitute(Alu.a.Bitor(*Alu.b))
		break
	case 7:
		Alu.s.Substitute(Alu.a.Bitxor(*Alu.b))
		break
	case 8:
		Alu.s.Substitute(Alu.a.And(*Alu.b))
		break
	case 9:
		Alu.s.Substitute(Alu.a.Or(*Alu.b))
		break
	case 10:
		Alu.s.Substitute(Alu.a.Equal(*Alu.b))
		break
	case 11:
		Alu.s.Substitute(Alu.a.NE(*Alu.b))
		break
	case 12:
		Alu.s.Substitute(Alu.x.GE(*Alu.y))
		break
	case 13:
		Alu.s.Substitute(Alu.x.LE(*Alu.y))
		break
	case 14:
		Alu.s.Substitute(Alu.x.GT(*Alu.y))
		break
	case 15:
		Alu.s.Substitute(Alu.x.LT(*Alu.y))
		break
	case 16:
		Alu.s.Substitute(Alu.a.Neg())
		break
	case 17:
		Alu.s.Substitute(Alu.a.Bnot())
		break
	case 18:
		Alu.s.Substitute(Alu.a.Not())
		break
	default:
		Alu.s.Substitute(*variable.CreateBits("16'hxxxx"))
		break
	}
}
