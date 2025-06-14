package main

import "fmt"

// マップ

func main() {

	// 文字列をキー、整数をバリューとするマップ
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make関数を使ってマップを作成
	m4 := make(map[int]string)
	fmt.Println(m4)

	// 要素の追加
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	// 要素の取得
	fmt.Println(m["A"])
	fmt.Println(m4[2])

	// エラーハンドリング
	// キーが存在しない場合、バリューの型のデフォルト値が返る
	_, ok := m4[1]
	fmt.Println(ok)
	s2, ok2 := m4[3]
	if !ok2 {
		fmt.Println("error")
	}
	fmt.Println(s2, ok2)

	// 要素の削除
	delete(m4, 1)
	fmt.Println(m4)

	var m5 = map[int]bool{1: true, 2: false, 3: false}
	fmt.Println(m5[1])

	m5[4] = true
	fmt.Println(m5[4])

	m6 := make(map[string]int)
	m6["Apple"] = 200
	m6["Banana"] = 300
	m6["Orange"] = 400
	fmt.Println(m6)

	delete(m6, "Banana")
	fmt.Println(m6)
}