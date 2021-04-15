package generated

import "github.com/verilog2Go/src/variable"

type elelock struct{
    clk, reset, close, tenkey, lock, ke1, ke2, key[0:3], match, key_enbl *variable.BitArray
}

func Elelock(args *elelock) elelock{
    args.clk.AddPosedgeObserver(args)
    args.reset.AddPosedgeObserver(args)
    args.clk.AddPosedgeObserver(args)
    args.reset.AddPosedgeObserver(args)
    args.clk.AddPosedgeObserver(args)
    args.reset.AddPosedgeObserver(args)
    return *args
}

func (elelock *elelock) Exec() {
    elelock.match.Assign(elelock.SECRET_3.Equal(*elelock.key.Get(3)).And(*elelock.SECRET_2.Equal(*elelock.key.Get(2)).And(*elelock.SECRET_1.Equal(*elelock.key.Get(1)).And(*elelock.SECRET_0.Equal(*elelock.key.Get(0))))))
    elelock.key_enbl.Assign(elelock.ke1.Bitand(*elelock.ke2.Not()))
}

func (elelock *elelock) PreAlways() []variable.BitArray{
    var1 := *variable.CreateBitArray(8, 0)
    var2 := *variable.CreateBitArray(8, 0)
    var3 := *variable.CreateBitArray(8, 0)
    var4 := *variable.CreateBitArray(8, 0)
    var5 := *variable.CreateBitArray(8, 0)
    var6 := *variable.CreateBitArray(8, 0)
    var7 := *variable.CreateBitArray(8, 0)
    var8 := *variable.CreateBitArray(8, 0)
    var9 := *variable.CreateBitArray(8, 0)
    var10 := *variable.CreateBitArray(8, 0)
    var11 := *variable.CreateBitArray(8, 0)
    var12 := *variable.CreateBitArray(8, 0)
    if variable.CreateBitArray(1, 1).Equal(*elelock.reset) {
        var1.Assign(variable.CreateBitArray(4, 15))
        var2.Assign(variable.CreateBitArray(4, 15))
        var3.Assign(variable.CreateBitArray(4, 15))
        var4.Assign(variable.CreateBitArray(4, 15))
    } else{
        if variable.CreateBitArray(1, 1).Equal(*elelock.close) {
            var5.Assign(variable.CreateBitArray(4, 15))
            var6.Assign(variable.CreateBitArray(4, 15))
            var7.Assign(variable.CreateBitArray(4, 15))
            var8.Assign(variable.CreateBitArray(4, 15))
        } else{
            if variable.CreateBitArray(1, 1).Equal(*elelock.key_enbl) {
                var9.Assign(*elelock.key.Get(2))
                var10.Assign(*elelock.key.Get(1))
                var11.Assign(*elelock.key.Get(0))
                var12.Assign(elelock.keyenc(elelock.tenkey))
            }
        }
    }
    return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7, var8, var9, var10, var11, var12}
}
func (elelock *elelock) PreAlways() []variable.BitArray{
    var1 := *variable.CreateBitArray(8, 0)
    var2 := *variable.CreateBitArray(8, 0)
    var3 := *variable.CreateBitArray(8, 0)
    var4 := *variable.CreateBitArray(8, 0)
    if variable.CreateBitArray(1, 1).Equal(*elelock.reset) {
        var1.Assign(*variable.CreateBitArray(1, 0))
        var2.Assign(*variable.CreateBitArray(1, 0))
    } else{
        var3.Assign(elelock.ke1)
        var4.Assign(elelock.tenkey.Reductionor())
    }
    return []variable.BitArray{var1, var2, var3, var4}
}
func (elelock *elelock) PreAlways() []variable.BitArray{
    var1 := *variable.CreateBitArray(8, 0)
    var2 := *variable.CreateBitArray(8, 0)
    var3 := *variable.CreateBitArray(8, 0)
    if variable.CreateBitArray(1, 1).Equal(*elelock.reset) {
        var1.Assign(*variable.CreateBitArray(1, 1))
    } else{
        if variable.CreateBitArray(1, 1).Equal(*elelock.close) {
            var2.Assign(*variable.CreateBitArray(1, 1))
        } else{
            if variable.CreateBitArray(1, 1).Equal(*elelock.match) {
                var3.Assign(*variable.CreateBitArray(1, 0))
            }
        }
    }
    return []variable.BitArray{var1, var2, var3}
}
func (elelock *elelock) Always(vars []variable.BitArray){
    if variable.CreateBitArray(1, 1).Equal(*elelock.reset) {
        elelock.key[3].Assign(vars[0])
        elelock.key[2].Assign(vars[1])
        elelock.key[1].Assign(vars[2])
        elelock.key[0].Assign(vars[3])
    } else{
        if variable.CreateBitArray(1, 1).Equal(*elelock.close) {
            elelock.key[3].Assign(vars[4])
            elelock.key[2].Assign(vars[5])
            elelock.key[1].Assign(vars[6])
            elelock.key[0].Assign(vars[7])
        } else{
            if variable.CreateBitArray(1, 1).Equal(*elelock.key_enbl) {
                elelock.key[3].Assign(vars[8])
                elelock.key[2].Assign(vars[9])
                elelock.key[1].Assign(vars[10])
                elelock.key[0].Assign(vars[11])
            }
        }
    }
    elelock.Exec()
}
func (elelock *elelock) Always(vars []variable.BitArray){
    if variable.CreateBitArray(1, 1).Equal(*elelock.reset) {
        elelock.ke2.Assign(vars[0])
        elelock.ke1.Assign(vars[1])
    } else{
        elelock.ke2.Assign(vars[2])
        elelock.ke1.Assign(vars[3])
    }
    elelock.Exec()
}
func (elelock *elelock) Always(vars []variable.BitArray){
    if variable.CreateBitArray(1, 1).Equal(*elelock.reset) {
        elelock.lock.Assign(vars[0])
    } else{
        if variable.CreateBitArray(1, 1).Equal(*elelock.close) {
            elelock.lock.Assign(vars[1])
        } else{
            if variable.CreateBitArray(1, 1).Equal(*elelock.match) {
                elelock.lock.Assign(vars[2])
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

