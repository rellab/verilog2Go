package variable

type PosedgeObserver interface {
	PreAlways()
	Always([]BitArray)
}

type NegedgeObserver interface {
	PreAlways()
	Always([]BitArray)
}

func (ba *BitArray) AddPosedgeObserver(po PosedgeObserver) {
	ba.pos = append(ba.pos, po)
}

func (ba *BitArray) AddNegedgeObserver(no NegedgeObserver) {
	ba.neg = append(ba.neg, no)
}

func (ba BitArray) NotifyPosedgeObserver() {
	// for i := len(ba.pos) - 1; i >= 0; i-- {
	// 	ba.pos[i].Always()
	// }

	// 全てのPreAlways終了後にAlwaysを実行する
}

func (ba BitArray) NotifyNegedgeObserver() {
	// ba.neg.Always()
	// for i := 0; i < len(ba.neg); i++ {
	// 	ba.neg[i].Always()
	// }
}
