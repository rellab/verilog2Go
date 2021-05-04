package generated

import "github.com/verilog2Go/src/variable"

type cnt_unit struct{
    CK, RES, EN, Q, CA *variable.BitArray
}

func Cnt_unit(args *cnt_unit) cnt_unit{
    args.CK.AddPosedgeObserver(args.PreAlways1, args.Always1)
    args.RES.AddPosedgeObserver(args.PreAlways1, args.Always1)
    return *args
}

func (cnt_unit *cnt_unit) Exec() {
    cnt_unit.CA.Assign(cnt_unit.Q.Bitand(*cnt_unit.EN))
}

func (cnt_unit *cnt_unit) PreAlways1() []variable.BitArray{
    var1 := *variable.CreateBitArray(0, cnt_unit.CK.ToInt())
    var2 := *variable.CreateBitArray(0, cnt_unit.RES.ToInt())
    var3 := *variable.CreateBitArray(0, cnt_unit.EN.ToInt())
    var4 := *variable.CreateBitArray(8, 0)
    var5 := *variable.CreateBitArray(8, 0)
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
        var4.Assign(*variable.CreateBitArray(1, 0))
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
            var5.Assign(cnt_unit.Q.Not())
        }
    }
    return []variable.BitArray{var1, var2, var3, var4, var5}
}
func (cnt_unit *cnt_unit) Always1(vars []variable.BitArray){
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
        cnt_unit.Q.Assign(vars[3])
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
            cnt_unit.Q.Assign(vars[4])
        }
    }
    cnt_unit.Exec()
}

