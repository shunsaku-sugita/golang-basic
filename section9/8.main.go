package main

import "fmt"

// channel
// 複数のゴルーチン間でデータの受け渡しを行うための仕組み
//　宣言、操作

func main() {
	// 指定しない場合は双方向のチャネルとなる
	var ch1 chan int

	// 受信専用のチャネル
	// var ch2 <-chan int

	// 送信専用のチャネル
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// バッファサイズを指定
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len",len(ch3))

	// チャネルの受信
	i := <-ch3
	fmt.Println(i)
	fmt.Println("len",len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len",len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len",len(ch3))

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	fmt.Println(<-ch3) // 過剰分を取り出して、デッドロックを回避する
	ch3 <- 6 // バッファサイズを超えるとブロックされ、デッドロックが発生する

	fmt.Println("len",len(ch3))

}