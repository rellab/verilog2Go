package generated

import "github.com/verilog2Go/src/variable"

type Counter struct {
	CLK, RES, Q *variable.BitArray
}

func NewCounter(args *Counter) Counter {
	args.CLK.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	args.RES.AddNegedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	return *args
}

func NewGoroutineCounter(in []chan int, out []chan int) *Counter {
	Counter := &Counter{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(4)}
	go Counter.run(in, out)
	return Counter
}

func (counter *Counter) Exec() {
}

func (counter *Counter) run(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				counter.CLK.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.Q.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				counter.RES.Set(v)
				bitArrays1 := counter.PreAlways1()
				counter.Always1(bitArrays1)
				counter.Exec()
				out[0] <- counter.Q.ToInt()
			} else {
				return
			}
		}
	}
}

func (Counter *Counter) PreAlways1() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Counter.CLK.ToInt())
	var2 := *variable.CreateBitArray(1, Counter.RES.ToInt())
	if variable.CheckBit(variable.CreateBitArray(1, 0).Equal(var2)) {
	} else {
		if variable.CheckBit(variable.CreateBitArray(4, 9).Equal(*Counter.Q)) {
		} else {
		}
	}
	return []variable.BitArray{var1, var2}
}
func (Counter *Counter) Always1(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBitArray(1, 0).Equal(vars[1])) {
	} else {
		if variable.CheckBit(variable.CreateBitArray(4, 9).Equal(*Counter.Q)) {
		} else {
		}
	}
}
