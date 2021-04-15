package generated

import (
	"github.com/verilog2Go/src/variable"
)

type cnt_unit struct {
	CK, RES, EN, Q, CA *variable.BitArray
}

func Cnt_unit(args *cnt_unit) cnt_unit {
	args.CK.AddPosedgeObserver(args)
	args.RES.AddPosedgeObserver(args)
	return *args
}

func (cnt_unit *cnt_unit) Exec() {
	cnt_unit.CA.Assign(cnt_unit.Q.Bitand(*cnt_unit.EN))
	// fmt.Printf("EN: %d\n", cnt_unit.EN.ToInt())
	// fmt.Printf("Q: %d\n", cnt_unit.Q.ToInt())
	// fmt.Printf("CA: %d\n", cnt_unit.CA.ToInt())
}

func (cnt_unit *cnt_unit) PreAlways() []variable.BitArray {
	var1 := *variable.CreateBitArray(8, 0)
	var2 := *variable.CreateBitArray(8, 0)
	var3 := *variable.CreateBitArray(8, cnt_unit.EN.ToInt())
	// fmt.Printf("EN: %d\n", cnt_unit.EN.ToInt())
	// var3 := *cnt_unit.EN
	if variable.CreateBitArray(1, 1).Equal(*cnt_unit.RES) {
		var1.Assign(*variable.CreateBitArray(1, 0))
	} else {
		if variable.CreateBitArray(1, 1).Equal(var3) {
			var2.Assign(cnt_unit.Q.Not())
		}
	}
	// fmt.Printf("var3: %d\n", var3.ToInt())
	return []variable.BitArray{var1, var2, var3}
}
func (cnt_unit *cnt_unit) Always(vars []variable.BitArray) {
	// fmt.Printf("vars[2]: %d\n", vars[2].ToInt())
	if variable.CreateBitArray(1, 1).Equal(*cnt_unit.RES) {
		cnt_unit.Q.Assign(vars[0])
	} else {
		if variable.CreateBitArray(1, 1).Equal(vars[2]) {
			cnt_unit.Q.Assign(vars[1])
		}
	}
	cnt_unit.Exec()
}
