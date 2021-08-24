package variable

type PosedgeObserver struct {
	PreAlways func() []BitArray
	Always    func(vars []BitArray)
	Exec      func()
}

type NegedgeObserver struct {
	PreAlways func() []BitArray
	Always    func(vars []BitArray)
	Exec      func()
}

func (ba *BitArray) AddPosedgeObserver(preAlways func() []BitArray, always func(vars []BitArray), exec func()) {
	ba.pos = append(ba.pos, PosedgeObserver{preAlways, always, exec})
}

func (ba *BitArray) AddNegedgeObserver(preAlways func() []BitArray, always func(vars []BitArray), exec func()) {
	ba.neg = append(ba.neg, NegedgeObserver{preAlways, always, exec})
}

func (ba BitArray) NotifyPosedgeObserver() {
	// 全てのPreAlways終了後にAlwaysを実行する
	var bitArrays = make([][]BitArray, len(ba.pos))
	for i := 0; i < len(ba.pos); i++ {
		bitArrays[i] = ba.pos[i].PreAlways()
	}
	// Alwaysを実行する
	for i := 0; i < len(ba.pos); i++ {
		ba.pos[i].Always(bitArrays[i])
		ba.pos[i].Exec()
	}
}

func (ba BitArray) NotifyNegedgeObserver() {
	// 全てのPreAlways終了後にAlwaysを実行する
	var bitArrays = make([][]BitArray, len(ba.neg))
	for i := 0; i < len(ba.neg); i++ {
		bitArrays[i] = ba.neg[i].PreAlways()
	}
	// fmt.Println("end")
	// Alwaysを実行する
	for i := len(ba.neg) - 1; i >= 0; i-- {
		ba.neg[i].Always(bitArrays[i])
	}
}
