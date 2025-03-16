package main

import (
	"fmt"
	"os"
)

// defer
// 関数の処理が終わった後に実行される処理

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

// deferはスタックで管理される
// 逆順で実行される
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	// 無名関数を使って複数のdeferを実行できる
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // リソースの解放に頻繁に使われる

	file.Write([]byte("Hello"))
}