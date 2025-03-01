package main

import (
	"fmt"
	"time"
)

// Hello World

func main() {
    fmt.Println("Hello World")
    fmt.Println("The time is", time.Now())

    // 変数
    // 明示的な定義
    // var 変数名 型 = 値
    var i int = 100
    fmt.Println(i)

    var s string = "Hello Go"
    fmt.Println(s)

    var t, f bool = true, false
    fmt.Println(t,f)

    var (
        i2 int = 200
        s2 string = "Golang"
    )
    fmt.Println(i2, s2)

    var i3 int // 初期値は0
    var s3 string // 初期値は""
    fmt.Println(i3, s3)

    i3 = 300
    s3 = "Go"
    fmt.Println(i3, s3)
    
    // 暗黙的な定義 (型を省略して定義)
    // 変数名 := 値
    i4 := 400
    fmt.Println(i4)

}