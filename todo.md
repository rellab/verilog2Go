# cnt_unit.vをGoに変換する

完成イメージ
```go
package generated

import "github.com/verilog2Go/src/builder/variable"

type cnt_unit struct{
    CK, RES, EN, Q, CA variable.BitArray
}

func Cnt_unit(CK variable.BitArray, RES variable.BitArray, EN variable.BitArray, Q variable.BitArray, CA variable.BitArray) *cnt_unit {
    p := new(cnt_unit)
    p.CK = CK
    p.RES = RES
    p.EN = EN
    p.Q = Q
    p.CA = CA
    p.CK.AddPosedgeObserver(p)
    p.RES.AddPosedgeObserver(p)
    return p
}

func (cnt_unit *cnt_unit) Exec() {
    cnt_unit.CA.Assign(cnt_unit.Q.Bitand(cnt_unit.EN))
}

func (cnt_unit cnt_unit) Always() {
    var var1 BitArray
    var var2 BitArray
    if cnt_unit.RES.ToInt() == 1 {
        var1.Set(0)
    } else {
        if cnt_unit.EN.ToInt() == 1{
            var2.Set(cnt_unit.Q.Not())
        }
    }
    if cnt_unit.RES.ToInt() == 1 {
        cnt_unit.Q.Set(var1)
    } else {
        if cnt_unit.EN.ToInt() == 1{
            cnt_unit.Q.Set(var2)
        }
    }
}
```