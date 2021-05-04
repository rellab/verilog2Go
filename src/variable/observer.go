package variable

type PosedgeObserver struct {
	PreAlways func() []BitArray
	Always    func(vars []BitArray)
}

type NegedgeObserver struct {
	PreAlways func() []BitArray
	Always    func(vars []BitArray)
}

func (ba *BitArray) AddPosedgeObserver(preAlways func() []BitArray, always func(vars []BitArray)) {
	ba.pos = append(ba.pos, PosedgeObserver{preAlways, always})
}

func (ba *BitArray) AddNegedgeObserver(preAlways func() []BitArray, always func(vars []BitArray)) {
	ba.neg = append(ba.neg, NegedgeObserver{preAlways, always})
}

func (ba BitArray) NotifyPosedgeObserver() {
	// 全てのPreAlways終了後にAlwaysを実行する
	var bitArrays = make([][]BitArray, len(ba.pos))
	for i := 0; i < len(ba.pos); i++ {
		bitArrays[i] = ba.pos[i].PreAlways()
	}
	// fmt.Println("end")
	// Alwaysを実行する
	for i := 0; i < len(ba.pos); i++ {
		// for i := len(ba.pos) - 1; i >= 0; i-- {
		ba.pos[i].Always(bitArrays[i])
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
