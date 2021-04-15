package generated

import (
	"github.com/verilog2Go/src/variable"
)

type saikoro struct {
	ck, reset, enable, lamp, cnt variable.BitArray
}

func Saikoro(ck *variable.BitArray, reset *variable.BitArray, enable *variable.BitArray, lamp *variable.BitArray, cnt *variable.BitArray) saikoro {
	p := new(saikoro)
	p.ck = *ck
	p.reset = *reset
	p.enable = *enable
	p.lamp = *lamp
	p.cnt = *cnt
	ck.AddPosedgeObserver(p)
	reset.AddPosedgeObserver(p)
	return *p
}

func (saikoro *saikoro) Exec() {
	saikoro.lamp.Assign(saikoro.dec(saikoro.cnt))
}

func (saikoro *saikoro) PreAlways() []variable.BitArray {
	var1 := *variable.CreateBitArray(8, 0)
	var2 := *variable.CreateBitArray(8, 0)
	var3 := *variable.CreateBitArray(8, 0)
	if variable.CreateBitArray(1, 1).Equal(saikoro.reset) {
		var1.Assign(*variable.CreateBitArray(3, 1))
	} else {
		if variable.CreateBitArray(1, 1).Equal(saikoro.enable) {
			if variable.CreateBitArray(3, 6).Equal(saikoro.cnt) {
				var2.Assign(*variable.CreateBitArray(3, 1))
			} else {
				var3.Assign(variable.CreateBitArray(3, 1).Add(saikoro.cnt))
			}
		}
	}
	return []variable.BitArray{var1, var2, var3}
}
func (saikoro *saikoro) Always(vars []variable.BitArray) {
	if variable.CreateBitArray(1, 1).Equal(saikoro.reset) {
		saikoro.cnt.Assign(vars[0])
	} else {
		if variable.CreateBitArray(1, 1).Equal(saikoro.enable) {
			if variable.CreateBitArray(3, 6).Equal(saikoro.cnt) {
				saikoro.cnt.Assign(vars[1])
			} else {
				saikoro.cnt.Assign(vars[2])
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
