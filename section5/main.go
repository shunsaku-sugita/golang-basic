package main

import "fmt"

// 定数

// 定数は変更できない
// 頭文字が大文字の定数は外部から参照できる

const Pi = 3.14

// 複数の定数を定義
const (
	URL = "http://example.com"
	SiteName = "TEST"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
	c0 = iota // 連続した整数を生成　
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	
	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}