package generated

import "github.com/verilog2Go/src/variable"

type Cnt_unit struct{
    CK, RES, EN, Q, CA *variable.BitArray
}

func NewCnt_unit(args *Cnt_unit) Cnt_unit{
    args.CK.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
    args.RES.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
    return *args
}

func NewGoroutineCnt_unit (in []chan int, out []chan int) *Cnt_unit{
    Cnt_unit := &Cnt_unit{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1)}
    go Cnt_unit.run(in, out)
    return Cnt_unit
}

func (cnt_unit *Cnt_unit) Exec() {
    cnt_unit.CA.Assign(cnt_unit.Q.Bitand(*cnt_unit.EN))
}

func (cnt_unit *Cnt_unit) run(in []chan int, out []chan int) {
    defer close(out[0])
    defer close(out[1])
    for {
        select {
        case v, ok := <-in[0]:
            if ok {
                cnt_unit.CK.Set(v)
                bitArrays1 := cnt_unit.PreAlways1()
                cnt_unit.Always1(bitArrays1)
                cnt_unit.Exec()
                out[0] <- cnt_unit.Q.ToInt()
                out[1] <- cnt_unit.CA.ToInt()
            } else {
                return 
            }
        case v, ok := <-in[1]:
            if ok {
                cnt_unit.RES.Set(v)
                bitArrays1 := cnt_unit.PreAlways1()
                cnt_unit.Always1(bitArrays1)
                cnt_unit.Exec()
                out[0] <- cnt_unit.Q.ToInt()
                out[1] <- cnt_unit.CA.ToInt()
            } else {
                return 
            }
        case v, ok := <-in[2]:
            if ok {
                cnt_unit.EN.Set(v)
                bitArrays1 := cnt_unit.PreAlways1()
                cnt_unit.Always1(bitArrays1)
                cnt_unit.Exec()
                out[0] <- cnt_unit.Q.ToInt()
                out[1] <- cnt_unit.CA.ToInt()
            } else {
                return 
            }
        }
    }
}

func (Cnt_unit *Cnt_unit) PreAlways1() []variable.BitArray{
    var1 := *variable.CreateBitArray(1, Cnt_unit.CK.ToInt())
    var2 := *variable.CreateBitArray(1, Cnt_unit.RES.ToInt())
    var3 := *variable.CreateBitArray(1, Cnt_unit.EN.ToInt())
    var4 := *variable.CreateBitArray(8, 0)
    var5 := *variable.CreateBitArray(8, 0)
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
        var4.Assign(*variable.CreateBitArray(1, 0))
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
            var5.Assign(Cnt_unit.Q.Not())
        }
    }
    return []variable.BitArray{var1, var2, var3, var4, var5}
}
func (Cnt_unit *Cnt_unit) Always1(vars []variable.BitArray){
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
        Cnt_unit.Q.Assign(vars[3])
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
            Cnt_unit.Q.Assign(vars[4])
        }
    }
}

