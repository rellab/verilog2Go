package variable

type PosedgeObserver interface {
	Always()
}

type NegedgeObserver interface {
	Always()
}

func (ba *BitArray) AddPosedgeObserver(po PosedgeObserver) {
	ba.pos = append(ba.pos, po)
}

func (ba *BitArray) AddNegedgeObserver(no NegedgeObserver) {
	ba.neg = append(ba.neg, no)
}

func (ba BitArray) NotifyPosedgeObserver() {
	// for pos := range ba.pos {
	// 	fmt.Println(pos)
	// }
	// for i := 0; i < len(ba.pos); i++ {
	// 	ba.pos[i].Always()
	// }
	for i := len(ba.pos) - 1; i >= 0; i-- {
		ba.pos[i].Always()
	}
}

func (ba BitArray) NotifyNegedgeObserver() {
	// ba.neg.Always()
	for i := 0; i < len(ba.neg); i++ {
		ba.neg[i].Always()
	}
}
