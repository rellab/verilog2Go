package variable

//Port はBitArrayを所有する
type Port struct {
	BitArray
}

//ToInt はBitArrayの持つ値を10進数で返す
func (port Port) ToInt() int {
	ret := 0
	length := len(port.bits)
	for i := length - 1; i >= 0; i-- {
		//bitの値がtrueの場合，対応する値を加算する
		if port.bits[i].value {
			ret |= 1 << i
		}
	}
	return ret
}

func (port *Port) Add(input Port) Port {
	result := Port{port.BitArray.Add(input.BitArray)}
	return result
}

func (port *Port) Sub(input Port) Port {
	result := Port{port.BitArray.Sub(input.BitArray)}
	return result
}

func (port *Port) Mul(input Port) Port {
	result := Port{port.BitArray.Mul(input.BitArray)}
	return result
}

func (port *Port) Assign(input Port) {
	port.BitArray.Assign(input.BitArray)
}
