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
	var m module1
	m.p.name = "parent"
	fmt.Println(m.p.getName())
}

func TestInterface(t *testing.T) {
	var m module1
	m.p.name = "parent"
	m.p.inter = m
	m.p.runAlways()
}
