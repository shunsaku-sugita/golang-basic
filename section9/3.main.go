package main

import "fmt"

// スライス
// copy

func main() {

	// スライスのコピー
	sl := []int{100, 200}
	sl2 := sl
	fmt.Println(sl, sl2)

	// スライスは参照型なので、sl2の値を変更するとslの値も変更される
	sl[0] = 1000
	fmt.Println(sl, sl2)

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)

	// copy関数（第一引数にコピー先、第二引数にコピー元）
	// copy関数はコピーした要素数を返す nにはコピーした要素数が入る
	n := copy(sl4, sl3)
	fmt.Println(n, sl4)
}