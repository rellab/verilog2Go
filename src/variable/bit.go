package variable

// Bit はビットを構造体で定義
type Bit struct {
	value bool
}

// BitArray はビットの配列を構造体で定義
type BitArray struct {
	bits []Bit
	pos  []PosedgeObserver
	neg  []NegedgeObserver
}

//InitBitArray はBitArrayを初期化する
func (ba *BitArray) InitBitArray(length int) {
	ba.bits = make([]Bit, length)
	for i := 0; i < length; i++ {
		bit := Bit{false}
		ba.bits[i] = bit
	}
}

func NewBitArray(length int) *BitArray {
	var ba BitArray
	ba.InitBitArray(length)
	return &ba
}

//Set はBitArrayのBitsに値をセットする
func (ba *BitArray) Set(value int) {
	preValue := ba.ToInt()
	length := len(ba.bits)
	comparison := 1 << length
	for i := 1; i <= length; i++ {
		ba.bits[length-i].value = (((value << i) & comparison) >> length) == 1
	}
	notify(*ba, preValue)
}

// CreateBitArray はvalueの値を持つBirArrayを返す
func CreateBitArray(length int, value int) *BitArray {
	var result BitArray
	result.InitBitArray(length + 1)
	result.Set(value)
	return &result
}

// Get はindexで指定したBitを持つBitArrayを返す
func (ba BitArray) Get(index int) *BitArray {
	var result BitArray
	result.InitBitArray(1)
	//スライスを代入
	result.bits = ba.bits[index : index+1]
	return &result
}

// Calc はvalueの値をもつBitArrayを返す
func (BitArray) Calc(value int, length int) BitArray {
	comparison := 1 << length
	var result BitArray
	result.InitBitArray(length)
	for i := 1; i <= length; i++ {
		result.bits[length-i].value = (((value << i) & comparison) >> length) == 1
	}
	return result
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

func notify(a BitArray, preValue int) {
	if a.ToInt() > preValue && a.pos != nil {
		a.NotifyPosedgeObserver()
	} else if a.ToInt() < preValue && a.neg != nil {
		a.NotifyNegedgeObserver()
	}
}

//Add はポート同士の加算を行う
func (ba BitArray) Add(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a+b, length)
	return result
}

//Sub はポート同士の減算を行う
func (ba BitArray) Sub(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a-b, length)
	return result
}

//Mul はポート同士の乗算を行う
func (ba BitArray) Mul(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a*b, length)
	return result
}

func (ba BitArray) And(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a == b {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

//Bitxor はポート同士の排他的論理和を返す
func (ba BitArray) Bitxor(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a^b, length)
	return result
}

//Bitand はポート同士の論理積を返す
func (ba BitArray) Bitand(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a&b, length)
	return result
}

//Bitor はポート同士の論理和を返す
func (ba BitArray) Bitor(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a|b, length)
	return result
}

func (ba *BitArray) Not() BitArray {
	length := len(ba.bits)
	var result BitArray
	result.InitBitArray(length)
	for i := 0; i < length; i++ {
		result.bits[i].value = !ba.bits[i].value
	}
	return result
}

func (ba *BitArray) Reductionor() BitArray {
	length := len(ba.bits)
	var result BitArray
	result.InitBitArray(1)
	for i := 0; i < length; i++ {
		if result.bits[i].value {
			result.Set(1)
		}
	}
	return result
}

//Equal はBitArrayの持つ値が等しいかどうかを返す
func (ba BitArray) Equal(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a == b {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

func CheckBit(input BitArray) bool {
	a := input.ToInt()
	if a > 0 {
		return true
	} else {
		return false
	}
}

// Assign は引数のBitArrayを割り当てる
func (ba *BitArray) Assign(result BitArray) {
	ba.Set(result.ToInt())
}
