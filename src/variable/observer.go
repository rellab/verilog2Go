package variable

type PosedgeObserver interface {
	Always()
}

type NegedgeObserver interface {
	Always()
}

func (ba *BitArray) AddPosedgeObserver(po PosedgeObserver) {
	ba.Pos = po
}

func (ba *BitArray) AddNegedgeObserver(no NegedgeObserver) {
	ba.neg = no
}

func (ba BitArray) NotifyPosedgeObserver() {
	ba.Pos.Always()
}

func (ba BitArray) NotifyNegedgeObserver() {
	ba.neg.Always()
}
