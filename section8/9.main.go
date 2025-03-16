package main

import (
	"fmt"
	"time"
)

// Go goroutine
// 並行処理

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub() // goをつけると並行処理になる
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}