package variable

// Bit はビットを構造体で定義
type Bit struct {
	value bool
}

// BitArray はビットの配列を構造体で定義
type BitArray struct {
	bits []Bit
}

//InitBit はBitを初期化する
func (b *Bit) InitBit() {
	b.value = false
}

//InitBitArray はBitArrayを初期化する
func (ba *BitArray) InitBitArray(length int) {
	for i := 0; i < length; i++ {
		bit := Bit{false}
		ba.bits = append(ba.bits, bit)
	}
}

// SetValue はBitのvalueに値をセットする
func (b *Bit) SetValue(value bool) {
	b.value = value
}

//Set はBitArrayのBitsに値をセットする
func (ba *BitArray) Set(value int) {
	length := len(ba.bits)
	comparison := 1 << length
	for i := 1; i <= length; i++ {
		ba.bits[length-i].SetValue((((value << i) & comparison) >> length) == 1)
	}
}

//ToInt はBitArrayの持つ値を10進数で返す
func (ba BitArray) ToInt() int {
	ret := 0
	length := len(ba.bits)
	for i := length - 1; i >= 0; i-- {
		//bitの値がtrueの場合，対応する値を加算する
		if ba.bits[i].value {
			ret |= 1 << i
		}
	}
	return ret
}

// Calc はvalueの値をもつBitArrayを返す
func (ba *BitArray) Calc(value int, length int) BitArray {
	comparison := 1 << length
	var result BitArray
	result.InitBitArray(length)
	for i := 1; i <= length; i++ {
		result.bits[length-i].SetValue((((value << i) & comparison) >> length) == 1)
	}
	return result
}

//Add はポート同士の加算を行う
func (ba *BitArray) Add(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a+b, length)
	return result
}

//Sub はポート同士の減算を行う
func (ba *BitArray) Sub(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a-b, length)
	return result
}

//Mul はポート同士の乗算を行う
func (ba *BitArray) Mul(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a*b, length)
	return result
}

func (ba *BitArray) Assign(result BitArray) {
	length := len(ba.bits)
	for i := 0; i < length; i++ {
		ba.bits[i].value = result.bits[i].value
	}
}
