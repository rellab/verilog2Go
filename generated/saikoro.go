package generated

import "github.com/verilog2Go/src/variable"

type Saikoro struct {
	ck, reset, enable, lamp, cnt *variable.BitArray
}

func NewSaikoro(args *Saikoro) Saikoro {
	args.ck.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	args.reset.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	return *args
}

func NewGoroutineSaikoro(in []chan int, out []chan int) *Saikoro {
	saikoro := &Saikoro{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(7), variable.NewBitArray(3)}
	go saikoro.run(in, out)
	return saikoro
}

func (saikoro *Saikoro) Exec() {
	saikoro.lamp.Assign(saikoro.dec(*saikoro.cnt))
}

func (saikoro *Saikoro) run(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				saikoro.ck.Set(v)
				bitArrays1 := saikoro.PreAlways1()
				saikoro.Always1(bitArrays1)
				saikoro.Exec()
				out[0] <- saikoro.lamp.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				saikoro.reset.Set(v)
				bitArrays1 := saikoro.PreAlways1()
				saikoro.Always1(bitArrays1)
				saikoro.Exec()
				out[0] <- saikoro.lamp.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				saikoro.enable.Set(v)
				bitArrays1 := saikoro.PreAlways1()
				saikoro.Always1(bitArrays1)
				saikoro.Exec()
				out[0] <- saikoro.lamp.ToInt()
			} else {
				return
			}
		}
	}
}

func (Saikoro *Saikoro) PreAlways1() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Saikoro.ck.ToInt())
	var2 := *variable.CreateBitArray(1, Saikoro.reset.ToInt())
	var3 := *variable.CreateBitArray(1, Saikoro.enable.ToInt())
	var4 := *variable.CreateBitArray(8, 0)
	var5 := *variable.CreateBitArray(8, 0)
	var6 := *variable.CreateBitArray(8, 0)
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
		var4.Assign(*variable.CreateBitArray(3, 1))
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
			if variable.CheckBit(variable.CreateBitArray(3, 6).Equal(*Saikoro.cnt)) {
				var5.Assign(*variable.CreateBitArray(3, 1))
			} else {
				var6.Assign(variable.CreateBitArray(3, 1).Add(*Saikoro.cnt))
			}
		}
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6}
}
func (Saikoro *Saikoro) Always1(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
		Saikoro.cnt.Assign(vars[3])
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
			if variable.CheckBit(variable.CreateBitArray(3, 6).Equal(*Saikoro.cnt)) {
				Saikoro.cnt.Assign(vars[4])
			} else {
				Saikoro.cnt.Assign(vars[5])
			}
		}
	}
}

func (Saikoro *Saikoro) dec(din variable.BitArray) variable.BitArray {
	dec := *variable.CreateBitArray(6, 0)
	switch din.ToInt() {
	case 1:
		dec.Set(8)
		break
	case 2:
		dec.Set(65)
		break
	case 3:
		dec.Set(28)
		break
	case 4:
		dec.Set(85)
		break
	case 5:
		dec.Set(93)
		break
	case 6:
		dec.Set(119)
		break
	}
	return dec
}
