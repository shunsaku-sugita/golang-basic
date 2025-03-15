package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 2
	if a == 2 {
		fmt.Println("Two")
	} else if a == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("I don't know")
	}

	// 短縮文
	if b := 100; b == 100 {
		fmt.Println("One Hundred")
	}
	// 注意点 (スコープ)
	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}