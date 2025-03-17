package main

import "fmt"

// スライス
// 可変長引数

// ...intのように引数の型の前に...をつけると、可変長引数を受け取ることができる
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// スライスを引数に渡すこともできる
	// スライスを引数に渡す場合は、スライスの後に...をつける
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(sl...))
}