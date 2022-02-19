package generated

import "github.com/verilog2Go/src/variable"

type Counter struct {
	clk, reset, load, inc, d, q *variable.BitArray
}

func NewCounter(args *Counter) Counter {
	args.clk.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	args.reset.AddNegedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	return *args
}

func NewGoroutineCounter(in []chan int, out []chan int) *Counter {
	counter := &Counter{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(16), variable.NewBitArray(16)}
	go counter.start(in, out)
	return counter
}

func (counter *Counter) Exec() {
}

func (counter *Counter) start(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				counter.clk.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.q.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				counter.reset.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.q.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				counter.load.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.q.ToInt()
			} else {
				return
			}
		case v, ok := <-in[3]:
			if ok {
				counter.inc.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.q.ToInt()
			} else {
				return
			}
		case v, ok := <-in[4]:
			if ok {
				counter.d.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.q.ToInt()
			} else {
				return
			}
		}
	}
}

func (Counter *Counter) PreAlways1() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Counter.clk.ToInt())
	var2 := *variable.CreateBitArray(1, Counter.reset.ToInt())
	var3 := *variable.CreateBitArray(1, Counter.load.ToInt())
	var4 := *variable.CreateBitArray(1, Counter.inc.ToInt())
	var5 := *variable.CreateBitArray(16, Counter.d.ToInt())
	var6 := *variable.CreateBitArray(8, 0)
	var7 := *variable.CreateBitArray(8, 0)
	var8 := *variable.CreateBitArray(8, 0)
	if variable.CheckBit(Counter.reset.Not()) {
		var6.Assign(*variable.CreateBits("16'h0000"))
	}
	if variable.CheckBit(*Counter.load) {
		var7.Assign(*Counter.d)
	}
	if variable.CheckBit(*Counter.inc) {
		var8.Assign(*Counter.q)
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7, var8}
}
func (Counter *Counter) Always1(vars []variable.BitArray) {
	if variable.CheckBit(Counter.reset.Not()) {
		Counter.q.Assign(vars[5])
	}
	if variable.CheckBit(*Counter.load) {
		Counter.q.Assign(vars[6])
	}
	if variable.CheckBit(*Counter.inc) {
		Counter.q.Assign(vars[7])
	}
}
