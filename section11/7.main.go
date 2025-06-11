package main

import "fmt"

// 独自型
type MyInt int

// 独自型のメソッドも定義できる
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// 独自型は基本型とは異なる型として扱われるため、基本型同士の演算はできない
	// i := 100
	// fmt.Println(mi + i)

	mi.Print()
}