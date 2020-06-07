package variable

// Bit はビットを構造体で定義
type Bit struct {
	value bool
}

// BitArray はビットの配列を構造体で定義
type BitArray struct {
	bits []Bit
}

//InitBitArray はBitArrayを初期化する
func (ba *BitArray) InitBitArray(value int, length int) BitArray {
	comparison := 1 << length
	for i := 1; i <= length; i++ {
		bit := Bit{(((value << i) & comparison) >> length) == 1}
		if len(ba.bits) == 0 {
			ba.bits = append(ba.bits, bit)
		} else {
			ba.bits, ba.bits[0] = append(ba.bits[:1], ba.bits[0:]...), bit
		}
	}
	return *ba
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

//Add はポート同士の加算を行う
func (ba BitArray) Add(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.InitBitArray(a+b, length)
	return result
}

//Sub はポート同士の減算を行う
func (ba BitArray) Sub(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.InitBitArray(a-b, length)
	return result
}

//Mul はポート同士の乗算を行う
func (ba BitArray) Mul(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.InitBitArray(a*b, length)
	return result
}

// Assign は引数のBitArrayを割り当てる
func (ba *BitArray) Assign(result BitArray) {
	length := len(ba.bits)
	for i := 0; i < length; i++ {
		ba.bits[i].value = result.bits[i].value
	}
}
