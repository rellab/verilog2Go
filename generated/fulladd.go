package generated

import "github.com/verilog2Go/src/variable"

type fulladd struct{
    A, B, CIN, Q, COUT variable.BitArray
}

func Fulladd(args *fulladd) fulladd{
    return *args
}

func (fulladd *fulladd) Exec() {
    fulladd.Q.Assign(fulladd.CIN.Bitxor(fulladd.B.Bitxor(fulladd.A)))
    fulladd.COUT.Assign(fulladd.A.Bitand(fulladd.CIN).Bitor(fulladd.CIN.Bitand(fulladd.B).Bitor(fulladd.B.Bitand(fulladd.A))))
}


