package variable

import (
	"strconv"
	"strings"
)

// Bit はビットを構造体で定義
type Bit struct {
	value         bool
	indefVal      bool
	highInpedance bool
}

// BitArray はビットの配列を構造体で定義
type BitArray struct {
	id     string
	testID string
	bits   []Bit
	pos    []PosedgeObserver
	neg    []NegedgeObserver
}

//InitBitArray はBitArrayを初期化する
func (ba *BitArray) InitBitArray(length int) {
	ba.bits = make([]Bit, length)
	for i := 0; i < length; i++ {
		bit := Bit{false, false, false}
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

func (ba *BitArray) SetBits(value string) {
	// value => ("4'b1111")
	// value holds the value of each bit as a string type
	values := strings.Split(value, "'")
	// length, _ := strconv.Atoi(values[0])
	var res int64
	var ok error
	str := values[1][1:]
	switch values[1][0] {
	case 'b':
		res, ok = strconv.ParseInt(str, 2, 64)
		break
	case 'o':
		res, ok = strconv.ParseInt(str, 8, 64)
		break
	case 'd':
		res, ok = strconv.ParseInt(str, 10, 64)
		break
	case 'h':
		res, ok = strconv.ParseInt(str, 16, 64)
		break
	default:
		break
	}
	if ok == nil {
		ba.Set(int(res))
	} else {
		ba.toIndef()
	}
}

func (ba *BitArray) SetId(id string) {
	ba.id = id
}

func (ba *BitArray) GetId() string {
	return ba.id
}

func (ba *BitArray) SetTestId(id string) {
	ba.testID = id
}

func (ba *BitArray) GetTestId() string {
	return ba.testID
}

func (ba *BitArray) GetBit(index int) Bit {
	return ba.bits[index]
}

func (ba *BitArray) GetBits() []Bit {
	return ba.bits
}

func CreateBits(value string) *BitArray {
	// value => ("4'b1111")
	// value holds the value of each bit as a string type
	values := strings.Split(value, "'")
	length, _ := strconv.Atoi(values[0])
	result := NewBitArray(length)
	var res int64
	var ok error
	str := values[1][1:]
	switch values[1][0] {
	case 'b':
		res, ok = strconv.ParseInt(str, 2, 64)
		break
	case 'o':
		res, ok = strconv.ParseInt(str, 8, 64)
		break
	case 'd':
		res, ok = strconv.ParseInt(str, 10, 64)
		break
	case 'h':
		res, ok = strconv.ParseInt(str, 16, 64)
		break
	default:
		break
	}
	if ok == nil {
		result.Set(int(res))
	} else {
		result.toIndef()
	}
	return result
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
	if a.ToInt() != preValue {
		writeSignal(a)
	}
}

func searchIndef(ba BitArray) bool {
	for _, v := range ba.bits {
		if v.indefVal {
			return true
		}
	}
	return false
}

func (ba *BitArray) toIndef() {
	for i, _ := range ba.bits {
		ba.bits[i].indefVal = true
	}
}

func (ba *BitArray) Substitute(input BitArray) {
	for i, _ := range ba.bits {
		ba.bits[i] = input.bits[i]
	}
}

//Add はポート同士の加算を行う
func (ba BitArray) Add(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a+b, length)
	}
	return result
}

//Sub はポート同士の減算を行う
func (ba BitArray) Sub(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a-b, length)
	}
	return result
}

//Mul はポート同士の乗算を行う
func (ba BitArray) Mul(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a*b, length)
	}
	return result
}

func (ba BitArray) SHL(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a<<b, length)
	}
	return result
}

func (ba BitArray) SHR(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a>>b, length)
	}
	return result
}

func (ba BitArray) And(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a > 0 && b > 0 {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

func (ba BitArray) Or(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a > 0 || b > 0 {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

//Bitxor はポート同士の排他的論理和を返す
func (ba BitArray) Bitxor(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a^b, length)
	}
	return result
}

//Bitand はポート同士の論理積を返す
func (ba BitArray) Bitand(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a&b, length)
	}
	return result
}

//Bitor はポート同士の論理和を返す
func (ba BitArray) Bitor(input BitArray) BitArray {
	var result BitArray
	result.InitBitArray(len(ba.bits))
	if searchIndef(ba) || searchIndef(input) {
		result.toIndef()
	} else {
		a := ba.ToInt()
		b := input.ToInt()
		length := len(ba.bits)
		result = result.Calc(a|b, length)
	}
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

func (ba *BitArray) Neg() BitArray {
	length := len(ba.bits)
	var result BitArray
	result.InitBitArray(length)
	for i := 0; i < length; i++ {
		result.bits[i].value = !ba.bits[i].value
	}
	return result
}

func (ba *BitArray) Bnot() BitArray {
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
		if ba.bits[i].value {
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

func (ba BitArray) NE(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a != b {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

func (ba BitArray) GE(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a >= b {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

func (ba BitArray) LE(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a <= b {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

func (ba BitArray) GT(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a > b {
		return *CreateBitArray(0, 1)
	} else {
		return *CreateBitArray(0, 0)
	}
}

func (ba BitArray) LT(input BitArray) BitArray {
	a := ba.ToInt()
	b := input.ToInt()
	if a < b {
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
