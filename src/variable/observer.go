package variable

import (
	"fmt"
	"sync"
)

type PosedgeObserver interface {
	PreAlways() []BitArray
	Always([]BitArray)
}

type NegedgeObserver interface {
	PreAlways() []BitArray
	Always([]BitArray)
}

func (ba *BitArray) AddPosedgeObserver(po PosedgeObserver) {
	ba.pos = append(ba.pos, po)
}

func (ba *BitArray) AddNegedgeObserver(no NegedgeObserver) {
	ba.neg = append(ba.neg, no)
}

func (ba BitArray) NotifyPosedgeObserver() {
	// 全てのPreAlways終了後にAlwaysを実行する
	var wg sync.WaitGroup
	var bitArrays = make([][]BitArray, len(ba.pos))
	for i := 0; i < len(ba.pos); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// fmt.Println("start")
			bitArrays[i] = ba.pos[i].PreAlways()
			// fmt.Printf("%d, %d\n", i, bitArrays[i][1].ToInt())
		}(i)
	}
	// 上記処理の終了を待つ
	wg.Wait()
	fmt.Println("end")
	// Alwaysを実行する
	for i := len(ba.pos) - 1; i >= 0; i-- {
		// ba.pos[i].Always(bitArrays[i])
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ba.pos[i].Always(bitArrays[i])
		}(i)
	}
	wg.Wait()
}

func (ba BitArray) NotifyNegedgeObserver() {
	// ba.neg.Always()
	// for i := 0; i < len(ba.neg); i++ {
	// 	ba.neg[i].Always()
	// }
}
