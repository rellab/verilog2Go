package generated

import "github.com/verilog2Go/src/variable"

type counter_unit struct {
	ck, res, q, ca *variable.BitArray
}

func Counter_unit(args *counter_unit) counter_unit {
	return *args
}

func (counter_unit *counter_unit) Exec() {
	cu0 := NewCnt_unit(&Cnt_unit{Q: counter_unit.q.Get(0), CA: counter_unit.ca.Get(0), CK: counter_unit.ck, RES: counter_unit.res, EN: variable.CreateBitArray(1, 1)})
	cu0.Exec()
	cu1 := NewCnt_unit(&Cnt_unit{RES: counter_unit.res, EN: counter_unit.ca.Get(0), Q: counter_unit.q.Get(1), CA: counter_unit.ca.Get(1), CK: counter_unit.ck})
	cu1.Exec()
	cu2 := NewCnt_unit(&Cnt_unit{CK: counter_unit.ck, RES: counter_unit.res, EN: counter_unit.ca.Get(1), Q: counter_unit.q.Get(2), CA: counter_unit.ca.Get(2)})
	cu2.Exec()
	cu3 := NewCnt_unit(&Cnt_unit{CK: counter_unit.ck, RES: counter_unit.res, EN: counter_unit.ca.Get(2), Q: counter_unit.q.Get(3), CA: counter_unit.ca.Get(3)})
	cu3.Exec()
}
