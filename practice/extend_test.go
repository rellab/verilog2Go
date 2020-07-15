package practice

import (
	"fmt"
	"testing"
)

func TestGetName(t *testing.T) {
	p := &parent{
		name: "test",
	}
	fmt.Println(p.getName())
}

func TestCount(t *testing.T) {
	var c child
	c.counter = 0
	c.p.name = "parent"
	c.count()
	fmt.Println(c.p.getName())
}
