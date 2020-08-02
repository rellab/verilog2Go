package practice

import "fmt"

type parent struct {
	name  string
	inter testInter
}

type module1 struct {
	p parent
}

type testInter interface {
	always()
}

func (p parent) getName() string {
	return p.name
}

func (p parent) runAlways() {
	p.inter.always()
}

// module1にalways()を持たせることでtestInterに代入できる
func (m module1) always() {
	fmt.Println("module1")
}

func main() {
	p := parent{name: "Taro"} // 値型の変数を用意する
	(&p).Greet("Hi")
}

func (p *parent) Greet(smg string) {
	fmt.Println(smg)
}
