package main

import "fmt"

// マップ
// for

func main() {
	m := map[string]int{
		"Apple": 100,
		"Banana": 200,
	}
	
	for k, v := range m {
		fmt.Println(k, v)
	}

	// キーだけ取り出す
	for k := range m {
		fmt.Println(k)
	}

	// バリューだけ取り出す
	for _, v := range m {
		fmt.Println(v)
	}
}