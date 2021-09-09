package generated

import "github.com/verilog2Go/src/variable"

type Elelock struct {
	clk, reset, close, tenkey, lock, ke1, ke2, match, key_enbl, SECRET_3, SECRET_2, SECRET_1, SECRET_0 *variable.BitArray
	key                                                                                                []*variable.BitArray
}

func NewElelock(args *Elelock) Elelock {
	args.clk.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	args.reset.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	args.clk.AddPosedgeObserver(args.PreAlways2, args.Always2, args.Exec)
	args.reset.AddPosedgeObserver(args.PreAlways2, args.Always2, args.Exec)
	args.clk.AddPosedgeObserver(args.PreAlways3, args.Always3, args.Exec)
	args.reset.AddPosedgeObserver(args.PreAlways3, args.Always3, args.Exec)
	args.SECRET_3 = variable.CreateBitArray(4, 5)
	args.SECRET_2 = variable.CreateBitArray(4, 9)
	args.SECRET_1 = variable.CreateBitArray(4, 6)
	args.SECRET_0 = variable.CreateBitArray(4, 3)
	return *args
}

func NewGoroutineElelock(in []chan int, out []chan int) *Elelock {
	Elelock := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(10), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(4), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), []*variable.BitArray{variable.NewBitArray(4), variable.NewBitArray(4), variable.NewBitArray(4), variable.NewBitArray(4)}}
	go Elelock.run(in, out)
	return Elelock
}

func (elelock *Elelock) Exec() {
	elelock.match.Assign(elelock.SECRET_3.Equal(*elelock.key[3]).And(elelock.SECRET_2.Equal(*elelock.key[2]).And(elelock.SECRET_1.Equal(*elelock.key[1]).And(elelock.SECRET_0.Equal(*elelock.key[0])))))
	elelock.key_enbl.Assign(elelock.ke1.Bitand(elelock.ke2.Not()))
}

func (elelock *Elelock) run(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				elelock.clk.Set(v)
				bitArrays1 := elelock.PreAlways1()
				bitArrays2 := elelock.PreAlways2()
				bitArrays3 := elelock.PreAlways3()
				elelock.Always1(bitArrays1)
				elelock.Always2(bitArrays2)
				elelock.Always3(bitArrays3)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				elelock.reset.Set(v)
				bitArrays1 := elelock.PreAlways1()
				bitArrays2 := elelock.PreAlways2()
				bitArrays3 := elelock.PreAlways3()
				elelock.Always1(bitArrays1)
				elelock.Always2(bitArrays2)
				elelock.Always3(bitArrays3)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				elelock.close.Set(v)
				bitArrays1 := elelock.PreAlways1()
				bitArrays2 := elelock.PreAlways2()
				bitArrays3 := elelock.PreAlways3()
				elelock.Always1(bitArrays1)
				elelock.Always2(bitArrays2)
				elelock.Always3(bitArrays3)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		case v, ok := <-in[3]:
			if ok {
				elelock.tenkey.Set(v)
				bitArrays1 := elelock.PreAlways1()
				bitArrays2 := elelock.PreAlways2()
				bitArrays3 := elelock.PreAlways3()
				elelock.Always1(bitArrays1)
				elelock.Always2(bitArrays2)
				elelock.Always3(bitArrays3)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		}
	}
}

func (Elelock *Elelock) PreAlways1() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Elelock.clk.ToInt())
	var2 := *variable.CreateBitArray(1, Elelock.reset.ToInt())
	var3 := *variable.CreateBitArray(1, Elelock.close.ToInt())
	var4 := *variable.CreateBitArray(10, Elelock.tenkey.ToInt())
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
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
			var9.Assign(*variable.CreateBitArray(4, 15))
			var10.Assign(*variable.CreateBitArray(4, 15))
			var11.Assign(*variable.CreateBitArray(4, 15))
			var12.Assign(*variable.CreateBitArray(4, 15))
		} else {
			if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*Elelock.key_enbl)) {
				var13.Assign(*Elelock.key[2])
				var14.Assign(*Elelock.key[1])
				var15.Assign(*Elelock.key[0])
				var16.Assign(Elelock.keyenc(*Elelock.tenkey))
			}
		}
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7, var8, var9, var10, var11, var12, var13, var14, var15, var16}
}
func (Elelock *Elelock) PreAlways2() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Elelock.clk.ToInt())
	var2 := *variable.CreateBitArray(1, Elelock.reset.ToInt())
	var3 := *variable.CreateBitArray(1, Elelock.close.ToInt())
	var4 := *variable.CreateBitArray(10, Elelock.tenkey.ToInt())
	var5 := *variable.CreateBitArray(8, 0)
	var6 := *variable.CreateBitArray(8, 0)
	var7 := *variable.CreateBitArray(8, 0)
	var8 := *variable.CreateBitArray(8, 0)
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
		var5.Assign(*variable.CreateBitArray(1, 0))
		var6.Assign(*variable.CreateBitArray(1, 0))
	} else {
		var7.Assign(*Elelock.ke1)
		var8.Assign(Elelock.tenkey.Reductionor())
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7, var8}
}
func (Elelock *Elelock) PreAlways3() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Elelock.clk.ToInt())
	var2 := *variable.CreateBitArray(1, Elelock.reset.ToInt())
	var3 := *variable.CreateBitArray(1, Elelock.close.ToInt())
	var4 := *variable.CreateBitArray(10, Elelock.tenkey.ToInt())
	var5 := *variable.CreateBitArray(8, 0)
	var6 := *variable.CreateBitArray(8, 0)
	var7 := *variable.CreateBitArray(8, 0)
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var2)) {
		var5.Assign(*variable.CreateBitArray(1, 1))
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(var3)) {
			var6.Assign(*variable.CreateBitArray(1, 1))
		} else {
			if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*Elelock.match)) {
				var7.Assign(*variable.CreateBitArray(1, 0))
			}
		}
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7}
}
func (Elelock *Elelock) Always1(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
		Elelock.key[3].Assign(vars[4])
		Elelock.key[2].Assign(vars[5])
		Elelock.key[1].Assign(vars[6])
		Elelock.key[0].Assign(vars[7])
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
			Elelock.key[3].Assign(vars[8])
			Elelock.key[2].Assign(vars[9])
			Elelock.key[1].Assign(vars[10])
			Elelock.key[0].Assign(vars[11])
		} else {
			if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*Elelock.key_enbl)) {
				Elelock.key[3].Assign(vars[12])
				Elelock.key[2].Assign(vars[13])
				Elelock.key[1].Assign(vars[14])
				Elelock.key[0].Assign(vars[15])
			}
		}
	}
}
func (Elelock *Elelock) Always2(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
		Elelock.ke2.Assign(vars[4])
		Elelock.ke1.Assign(vars[5])
	} else {
		Elelock.ke2.Assign(vars[6])
		Elelock.ke1.Assign(vars[7])
	}
}
func (Elelock *Elelock) Always3(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[1])) {
		Elelock.lock.Assign(vars[4])
	} else {
		if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(vars[2])) {
			Elelock.lock.Assign(vars[5])
		} else {
			if variable.CheckBit(variable.CreateBitArray(1, 1).Equal(*Elelock.match)) {
				Elelock.lock.Assign(vars[6])
			}
		}
	}
}

func (Elelock *Elelock) keyenc(sw variable.BitArray) variable.BitArray {
	keyenc := *variable.CreateBitArray(3, 0)
	switch sw.ToInt() {
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
