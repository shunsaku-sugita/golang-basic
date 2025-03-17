package main

import "fmt"

// スライス
// append, make len, cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300) // 末尾に要素を追加
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	// makeはスライスを作成する関数
	// 第一引数はスライスの型、第二引数はスライスの長さ、第三引数はスライスの容量
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	// lenはスライスの長さを返す関数
	fmt.Println(len(sl2))

	// capはスライスの容量を返す関数
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(cap(sl3))
} 