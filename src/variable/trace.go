package variable

import (
	"strconv"
)

func Dump(time int) {
	source += "#" + strconv.Itoa(time) + "\n"
	if time == 0 {
		for _, v := range bitArrays {
			writeSignal(*v)
		}
	}
}

func writeSignal(ba BitArray) {
	if ba.GetTestId() != "" {
		source += showBynary(ba) + " " + ba.GetTestId() + "\n"
	}
}

func showBynary(ba BitArray) string {
	var bit string
	for _, v := range ba.GetBits() {
		if v.value {
			bit = "1" + bit
		} else {
			bit = "0" + bit
		}
	}
	return "b" + bit
}
