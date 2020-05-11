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

//Set はBitArrayのBitsに値をセットする
func (ba *BitArray) Set(value int) {
	var length = len(ba.bits)
	var comparison = 1 << length
	for i := 1; i <= length; i++ {
		// if (((value << i) & comparison) >> length) == 1 {
		// 	ba.bits[length - i].SetValue(true)
		// } else{
		// 	ba.bits[length - i].SetValue(false)
		// }
		ba.bits[length-i].SetValue((((value << i) & comparison) >> length) == 1)
	}
}
