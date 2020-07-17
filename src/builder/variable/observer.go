package variable

type PosedgeObserver interface {
	Always()
}

type NegedgeObserver interface {
	Always()
}

func (ba *BitArray) AddPosedgeObserver(po PosedgeObserver) {
	ba.pos = po
}

func (ba *BitArray) AddNegedgeObserver(no NegedgeObserver) {
	ba.neg = no
}

func (ba BitArray) NotifyPosedgeObserver() {
	ba.pos.Always()
}

func (ba BitArray) NotifyNegedgeObserver() {
	ba.neg.Always()
}
