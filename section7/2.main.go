package main

import "fmt"

func main() {
	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 2)
	fmt.Println(i5)

	// ========================================
	// 即時関数 (関数の定義と同時に実行)
	i6 := func(x, y int) int {
		return x + y
	}(1, 2) // ここで即時実行
	fmt.Println(i6)
}