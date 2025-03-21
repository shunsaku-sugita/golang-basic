package main

import "fmt"

// 関数
// ジェネレーターの実装

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}