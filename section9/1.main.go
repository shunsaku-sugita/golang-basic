package main

import "fmt"

// スライス
// 宣言、操作
// 配列と違い、要素数を指定しない
func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}	// スライス
	var arr [2]int = [2]int{100, 200} // 配列
	fmt.Println(arr, sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 値の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	// 3番目から4番目までの要素を取得
	fmt.Println(sl5[2:4])

	// 2番目までの要素を取得
	fmt.Println(sl5[:2])

	// 3番目から最後までの要素を取得
	fmt.Println(sl5[2:])

}