package variable

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var mark = []string{"#", "$", "%", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}

var moduleName string
var source string
var bitArrays []*BitArray

func Trace(module string, ba []*BitArray) {
	moduleName = module
	bitArrays = ba
	source = "$version Generated by verigo $end\n"
	source += "$date " + setDate() + "$end\n"
	source += "$timescale 1ns $end\n\n"

	source += "$scope module TOP $end\n"
	var signals string
	for i, v := range bitArrays {
		v.SetTestId(mark[i])
		signals += inputIndent(1) + "$var wire " + strconv.Itoa(len(v.GetBits())) + " " + v.GetTestId() + " " + v.GetId() + " [" + strconv.Itoa(len(v.GetBits())-1) + ":0] $end\n"
	}
	source += signals
	source += "$scope module " + moduleName + " $end\n"
	source += signals
	source += "$upscope $end\n"
	source += "$upscope $end\n"
	source += "enddefinitions $end\n\n"
}

func Close() {
	write()
}

func write() {
	file, err := os.Create(moduleName + ".vcd")
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, source) //書き込み
}

func setDate() string {
	var date string
	t := time.Now()

	switch t.Weekday() {
	case 0:
		date = "Sun "
	case 1:
		date = "Mon "
	case 2:
		date = "Tue "
	case 3:
		date = "Wed "
	case 4:
		date = "Thu "
	case 5:
		date = "Fri "
	case 6:
		date = "Sat "
	}
	switch t.Month() {
	case 1:
		date += "Jan "
	case 2:
		date += "Feb "
	case 3:
		date += "Mar "
	case 4:
		date += "Apr "
	case 5:
		date += "May "
	case 6:
		date += "Jun "
	case 7:
		date += "Jul "
	case 8:
		date += "Aug "
	case 9:
		date += "Sep "
	case 10:
		date += "Oct "
	case 11:
		date += "Nov "
	case 12:
		date += "Dec "
	}

	date += strconv.Itoa(t.Day()) + " "
	date += strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second()) + " "
	date += strconv.Itoa(t.Year()) + " "

	return date
}

func inputIndent(indent int) string {
	var result string
	for i := 0; i < indent*2; i++ {
		result += " "
	}
	return result
}
