package variable

// Bit はビットを構造体で定義
type Bit struct {
	value bool
	index int
}

// BitArray はビットの配列を構造体で定義
type BitArray struct {
	bits []Bit
}

//InitializeBit はBitを初期化する
func (b *Bit) InitializeBit(index int) {
	b.value = false
	b.index = index
}

//InitializeBitArray はBitArrayを初期化する
func (ba *BitArray) InitializeBitArray(length int) {
	for i := 0; i < length; i++ {
		var bit Bit
		bit.InitializeBit(i)
		ba.bits = append(ba.bits, bit)
	}
}

// SetValue はBitのvalueに値をセットする
func (b *Bit) SetValue(value bool) {
	b.value = value
}
