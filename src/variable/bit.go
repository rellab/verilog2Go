package variable

import "fmt"

// Bit はビットを構造体で定義
type Bit struct {
	value bool
}

// BitArray はビットの配列を構造体で定義
type BitArray struct {
	bits []Bit
	Pos  PosedgeObserver
	neg  NegedgeObserver
}

//InitBitArray はBitArrayを初期化する
func (ba *BitArray) InitBitArray(length int) {
	ba.bits = make([]Bit, length)
	for i := 0; i < length; i++ {
		bit := Bit{false}
		ba.bits[i] = bit
	}
}

//Set はBitArrayのBitsに値をセットする
func (ba *BitArray) Set(value int) {
	notify(*ba, value)
	length := len(ba.bits)
	comparison := 1 << length
	for i := 1; i <= length; i++ {
		ba.bits[length-i].value = (((value << i) & comparison) >> length) == 1
	}
}

// CreateBitArray はvalueの値を持つBirArrayを返す
func CreateBitArray(length int, value int) BitArray {
	var result BitArray
	result.InitBitArray(length)
	result.Set(value)
	return result
}

// Get はindexで指定したBitを持つBitArrayを返す
func (ba BitArray) Get(index int) BitArray {
	var result BitArray
	result.InitBitArray(1)
	//スライスを代入
	result.bits = ba.bits[index : index+1]
	return result
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

func notify(a BitArray, b int) {
	if a.ToInt() < b && a.Pos != nil {
		fmt.Println("notify posedge observer")
		a.NotifyPosedgeObserver()
	} else if a.ToInt() > b && a.neg != nil {
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
	notify(ba, result.ToInt())
	return result
}

//Sub はポート同士の減算を行う
func (ba BitArray) Sub(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a-b, length)
	notify(ba, result.ToInt())
	return result
}

//Mul はポート同士の乗算を行う
func (ba BitArray) Mul(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a*b, length)
	notify(ba, result.ToInt())
	return result
}

//Bitxor はポート同士の排他的論理和を返す
func (ba BitArray) Bitxor(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a^b, length)
	notify(ba, result.ToInt())
	return result
}

//Bitand はポート同士の論理積を返す
func (ba BitArray) Bitand(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a&b, length)
	notify(ba, result.ToInt())
	return result
}

//Bitor はポート同士の論理和を返す
func (ba BitArray) Bitor(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	length := len(ba.bits)
	var result BitArray
	result = result.Calc(a|b, length)
	notify(ba, result.ToInt())
	return result
}

func (ba BitArray) Not() BitArray {
	length := len(ba.bits)
	var result BitArray
	result.InitBitArray(length)
	for i := 0; i < length; i++ {
		result.bits[i].value = !ba.bits[i].value
	}
	notify(ba, result.ToInt())
	return result
}

//Equal はBitArrayの持つ値が等しいかどうかを返す
func (ba BitArray) Equal(input BitArray) bool {
	a := ba.ToInt()
	b := input.ToInt()
	if a == b {
		return true
	} else {
		return false
	}
}

// Assign は引数のBitArrayを割り当てる
func (ba *BitArray) Assign(result BitArray) {
	length := len(ba.bits)
	for i := 0; i < length; i++ {
		ba.bits[i].value = result.bits[i].value
	}
}
