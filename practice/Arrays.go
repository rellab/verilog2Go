package main

import "fmt"

func main() {
	//Arrayの宣言
	a := [3]int{1, 2, 3}
	//Sliceの宣言
	s := a[1:]

	fmt.Println("Array:", a)
	fmt.Println("Slice:", s)

	a[1] = 10
	fmt.Println("Array:", a)
	fmt.Println("Slice:", s)

	s[1] = 20
	fmt.Println("Array:", a)
	fmt.Println("Slice:", s)
	//sliceはポインタで参照している？

	b := a
	fmt.Println("Array:", a)
	fmt.Println("Substituted Array:", b)

	b[0] = 30
	fmt.Println("Array:", a)
	fmt.Println("Substituted Array:", b)
	//Array: [1 10 20]
	//Substituted Array: [30 10 20]
	//代入元の配列は変更されていない

	t := s
	t[0] = 50
	fmt.Println("Array:", a)
	fmt.Println("Slice:", s)
	fmt.Println("Substituted Slice:", t)
	// Array: [1 50 20]
	// Slice: [50 20]
	// Substituted Slice: [50 20]
	// sliceは参照渡し

	u := s[1]
	u = 0
	fmt.Println("Array:", a)
	fmt.Println("Slice:", s)
	fmt.Println("Content:", u)
	// Array: [1 50 20]
	// Slice: [50 20]
	// Content: 0
	// 要素は値渡し
}
