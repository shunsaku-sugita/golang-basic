package main

import "fmt"

// init
// init関数は、パッケージが初期化される際に実行される関数。
// パッケージが初期化される際は、main関数が実行される前に実行される。
func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Main")
}