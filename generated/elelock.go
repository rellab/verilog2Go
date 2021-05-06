package generated

import "github.com/verilog2Go/src/variable"

type elelock struct{
    clk, reset, close, tenkey, lock, ke1, ke2, match, key_enbl, SECRET_3, SECRET_2, SECRET_1, SECRET_0 *variable.BitArray
    key []*variable.BitArray
}

func Elelock(args *elelock) elelock{
    args.clk.AddPosedgeObserver(args.PreAlways1, args.Always1)
    args.reset.AddPosedgeObserver(args.PreAlways1, args.Always1)
    args.clk.AddPosedgeObserver(args.PreAlways2, args.Always2)
    args.reset.AddPosedgeObserver(args.PreAlways2, args.Always2)
    args.clk.AddPosedgeObserver(args.PreAlways3, args.Always3)
    args.reset.AddPosedgeObserver(args.PreAlways3, args.Always3)
    args.SECRET_3 = variable.CreateBitArray(4, 5)
    args.SECRET_2 = variable.CreateBitArray(4, 9)
    args.SECRET_1 = variable.CreateBitArray(4, 6)
    args.SECRET_0 = variable.CreateBitArray(4, 3)
    return *args
}

func (elelock *elelock) Exec() {
    elelock.match.Assign(elelock.SECRET_3.Equal(*elelock.key[3]).And(elelock.SECRET_2.Equal(*elelock.key[2]).And(elelock.SECRET_1.Equal(*elelock.key[1]).And(elelock.SECRET_0.Equal(*elelock.key[0])))))
    elelock.key_enbl.Assign(elelock.ke1.Bitand(elelock.ke2.Not()))
}

func (elelock *elelock) PreAlways1() []variable.BitArray{
    var1 := *variable.CreateBitArray(0, elelock.clk.ToInt())
    var2 := *variable.CreateBitArray(0, elelock.reset.ToInt())
    var3 := *variable.CreateBitArray(0, elelock.close.ToInt())
    var4 := *variable.CreateBitArray(9, elelock.tenkey.ToInt())
    var5 := *variable.CreateBitArray(8, 0)
    var6 := *variable.CreateBitArray(8, 0)
    var7 := *variable.CreateBitArray(8, 0)
    var8 := *variable.CreateBitArray(8, 0)
    var9 := *variable.CreateBitArray(8, 0)
    var10 := *variable.CreateBitArray(8, 0)
    var11 := *variable.CreateBitArray(8, 0)
    var12 := *variable.CreateBitArray(8, 0)
    var13 := *variable.CreateBitArray(8, 0)
    var14 := *variable.CreateBitArray(8, 0)
    var15 := *variable.CreateBitArray(8, 0)
    var16 := *variable.CreateBitArray(8, 0)
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
        var5.Assign(*variable.CreateBitArray(4, 15))
        var6.Assign(*variable.CreateBitArray(4, 15))
        var7.Assign(*variable.CreateBitArray(4, 15))
        var8.Assign(*variable.CreateBitArray(4, 15))
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
            var9.Assign(*variable.CreateBitArray(4, 15))
            var10.Assign(*variable.CreateBitArray(4, 15))
            var11.Assign(*variable.CreateBitArray(4, 15))
            var12.Assign(*variable.CreateBitArray(4, 15))
        } else{
            if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*elelock.key_enbl)) {
                var13.Assign(*elelock.key[2])
                var14.Assign(*elelock.key[1])
                var15.Assign(*elelock.key[0])
                var16.Assign(elelock.keyenc(*elelock.tenkey))
            }
        }
    }
    return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7, var8, var9, var10, var11, var12, var13, var14, var15, var16}
}
func (elelock *elelock) PreAlways2() []variable.BitArray{
    var1 := *variable.CreateBitArray(0, elelock.clk.ToInt())
    var2 := *variable.CreateBitArray(0, elelock.reset.ToInt())
    var3 := *variable.CreateBitArray(0, elelock.close.ToInt())
    var4 := *variable.CreateBitArray(9, elelock.tenkey.ToInt())
    var5 := *variable.CreateBitArray(8, 0)
    var6 := *variable.CreateBitArray(8, 0)
    var7 := *variable.CreateBitArray(8, 0)
    var8 := *variable.CreateBitArray(8, 0)
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
        var5.Assign(*variable.CreateBitArray(1, 0))
        var6.Assign(*variable.CreateBitArray(1, 0))
    } else{
        var7.Assign(*elelock.ke1)
        var8.Assign(elelock.tenkey.Reductionor())
    }
    return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7, var8}
}
func (elelock *elelock) PreAlways3() []variable.BitArray{
    var1 := *variable.CreateBitArray(0, elelock.clk.ToInt())
    var2 := *variable.CreateBitArray(0, elelock.reset.ToInt())
    var3 := *variable.CreateBitArray(0, elelock.close.ToInt())
    var4 := *variable.CreateBitArray(9, elelock.tenkey.ToInt())
    var5 := *variable.CreateBitArray(8, 0)
    var6 := *variable.CreateBitArray(8, 0)
    var7 := *variable.CreateBitArray(8, 0)
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
        var5.Assign(*variable.CreateBitArray(1, 1))
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
            var6.Assign(*variable.CreateBitArray(1, 1))
        } else{
            if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*elelock.match)) {
                var7.Assign(*variable.CreateBitArray(1, 0))
            }
        }
    }
    return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7}
}
func (elelock *elelock) Always1(vars []variable.BitArray){
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
        elelock.key[3].Assign(vars[4])
        elelock.key[2].Assign(vars[5])
        elelock.key[1].Assign(vars[6])
        elelock.key[0].Assign(vars[7])
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
            elelock.key[3].Assign(vars[8])
            elelock.key[2].Assign(vars[9])
            elelock.key[1].Assign(vars[10])
            elelock.key[0].Assign(vars[11])
        } else{
            if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*elelock.key_enbl)) {
                elelock.key[3].Assign(vars[12])
                elelock.key[2].Assign(vars[13])
                elelock.key[1].Assign(vars[14])
                elelock.key[0].Assign(vars[15])
            }
        }
    }
    elelock.Exec()
}
func (elelock *elelock) Always2(vars []variable.BitArray){
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
        elelock.ke2.Assign(vars[4])
        elelock.ke1.Assign(vars[5])
    } else{
        elelock.ke2.Assign(vars[6])
        elelock.ke1.Assign(vars[7])
    }
    elelock.Exec()
}
func (elelock *elelock) Always3(vars []variable.BitArray){
    if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
        elelock.lock.Assign(vars[4])
    } else{
        if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
            elelock.lock.Assign(vars[5])
        } else{
            if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*elelock.match)) {
                elelock.lock.Assign(vars[6])
            }
        }
    }
    elelock.Exec()
}

func (elelock *elelock) keyenc(sw variable.BitArray) variable.BitArray {
    keyenc := *variable.CreateBitArray(3, 0)
    switch sw.ToInt(){
    case 1:
        keyenc.Set(0)
        break
    case 2:
        keyenc.Set(1)
        break
    case 4:
        keyenc.Set(2)
        break
    case 8:
        keyenc.Set(3)
        break
    case 16:
        keyenc.Set(4)
        break
    case 32:
        keyenc.Set(5)
        break
    case 64:
        keyenc.Set(6)
        break
    case 128:
        keyenc.Set(7)
        break
    case 256:
        keyenc.Set(8)
        break
    case 512:
        keyenc.Set(9)
        break
    }
    return keyenc
}

