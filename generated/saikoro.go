package generated

import "github.com/verilog2Go/src/variable"

type saikoro struct {
	ck, reset, enable, lamp, cnt *variable.BitArray
}

func Saikoro(args *saikoro) saikoro {
	args.ck.AddPosedgeObserver(args)
	args.reset.AddPosedgeObserver(args)
	return *args
}

func (saikoro *saikoro) Exec() {
	saikoro.lamp.Assign(saikoro.dec(*saikoro.cnt))
}

func (saikoro *saikoro) PreAlways() []variable.BitArray {
	var1 := *variable.CreateBitArray(0, saikoro.ck.ToInt())
	var2 := *variable.CreateBitArray(0, saikoro.reset.ToInt())
	var3 := *variable.CreateBitArray(0, saikoro.enable.ToInt())
	var4 := *variable.CreateBitArray(8, 0)
	var5 := *variable.CreateBitArray(8, 0)
	var6 := *variable.CreateBitArray(8, 0)
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
		var4.Assign(*variable.CreateBitArray(3, 1))
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
			if variable.CheckBit(variable.CreateBitArray(3, 6).Equal(*saikoro.cnt)) {
				var5.Assign(*variable.CreateBitArray(3, 1))
			} else {
				var6.Assign(variable.CreateBitArray(3, 1).Add(*saikoro.cnt))
			}
		}
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6}
}
func (saikoro *saikoro) Always(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
		saikoro.cnt.Assign(vars[3])
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
			if variable.CheckBit(variable.CreateBitArray(3, 6).Equal(*saikoro.cnt)) {
				saikoro.cnt.Assign(vars[4])
			} else {
				saikoro.cnt.Assign(vars[5])
			}
		}
	}
	saikoro.Exec()
}

func (saikoro *saikoro) dec(din variable.BitArray) variable.BitArray {
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
