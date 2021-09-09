package generated

// import "github.com/verilog2Go/src/variable"

// type Counter_unit struct{
//     ck, res, q, ca *variable.BitArray
// }

// func NewCounter_unit(args *Counter_unit) Counter_unit{
//     return *args
// }

// func NewGoroutineCounter_unit (in []chan int, out []chan int) *Counter_unit{
//     Counter_unit := &Counter_unit{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(4), variable.NewBitArray(4)}
//     go Counter_unit.run(in, out)
//     return Counter_unit
// }

// func (counter_unit *Counter_unit) Exec() {
//     cu0 := Cnt_unit(Cnt_unit{CK : Counter_unit.ck, RES : Counter_unit.res, EN : variable.CreateBitArray(1, 1), Q : Counter_unit.q.Get(0), CA : Counter_unit.ca.Get(0)})
//     cu0.Exec()
//     cu1 := Cnt_unit(&cnt_unit{CK : Counter_unit.ck, RES : Counter_unit.res, EN : Counter_unit.ca.Get(0), Q : Counter_unit.q.Get(1), CA : Counter_unit.ca.Get(1)})
//     cu1.Exec()
//     cu2 := Cnt_unit(&cnt_unit{CK : Counter_unit.ck, RES : Counter_unit.res, EN : Counter_unit.ca.Get(1), Q : Counter_unit.q.Get(2), CA : Counter_unit.ca.Get(2)})
//     cu2.Exec()
//     cu3 := Cnt_unit(&cnt_unit{CK : Counter_unit.ck, RES : Counter_unit.res, EN : Counter_unit.ca.Get(2), Q : Counter_unit.q.Get(3), CA : Counter_unit.ca.Get(3)})
//     cu3.Exec()
// }

// func (counter_unit *Counter_unit) run(in []chan int, out []chan int) {
//     defer close(out[0])
//     for {
//         select {
//         case v, ok := <-in[0]:
//             if ok {
//                 counter_unit.ck.Set(v)
//                 counter_unit.Exec()
//                 out[0] <- counter_unit.q.ToInt()
//             } else {
//                 return
//             }
//         case v, ok := <-in[1]:
//             if ok {
//                 counter_unit.res.Set(v)
//                 counter_unit.Exec()
//                 out[0] <- counter_unit.q.ToInt()
//             } else {
//                 return
//             }
//         }
//     }
// }
