package generated
import "github.com/verilog2Go/src/variable"
type Adder struct{
a, b, q *variable.BitArray
}
func NewAdder(args *Adder) Adder{
return *args
}

func NewGoroutineAdder (in []chan int, out []chan int) *Adder{
adder := &Adder{variable.NewBitArray(4), variable.NewBitArray(4), variable.NewBitArray(4)}
go adder.run(in, out)
return adder
}

func (adder *Adder) Exec() {
adder.q.Assign(adder.b.Add(*adder.a))
}

func (adder *Adder) run(in []chan int, out []chan int) {
defer close(out[0])
for {
select {
case v, ok := <-in[0]:
if ok {
adder.a.Set(v)
adder.Exec()
out[0] <- adder.q.ToInt()
} else {
return 
}
case v, ok := <-in[1]:
if ok {
adder.b.Set(v)
adder.Exec()
out[0] <- adder.q.ToInt()
} else {
return 
}
}
}
}

