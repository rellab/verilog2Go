package practice

type parent struct {
	name string
}

type child struct {
	counter int
	p       parent
}

func (p parent) getName() string {
	return p.name
}

func childCount(c child) {
	c.count()
}

func (c child) count() {
	c.counter++
}
