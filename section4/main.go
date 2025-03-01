package main

import (
	"fmt"
	"strconv"
)



func main() {
	// 型
	// 整数型
	var i int = 100

	fmt.Println(i + 50)

	var i2 int64 = 200
	fmt.Printf("%T\n", i2) // %Tで型を表示

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

	// 浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T\n", fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	// 論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 文字列型の連結
	fmt.Println("Hello" + "World")

	// ダブルクォーテーションを表示する
	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0]) // バイト型で表示
	fmt.Println(string(s[0])) // 文字列型で表示

	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))
	
	// 配列型 
	// 要素数を後から変更できない
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Println(arr3[1])

	// 要素数を自動で
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 値の取り出し
	fmt.Println(arr2[0])

	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数の確認
	fmt.Println(len(arr1))


	// Interface型
	var x interface{}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = "Hello Interface"
	fmt.Println(x)
	x = [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)


	// 型変換
	var i6 int = 1
	var fl64_2 float64 = float64(i6)
	fmt.Printf("i = %T\n", i6)	
	fmt.Printf("fl64_2 = %T\n", fl64_2)

	var s1 string = "100"
	fmt.Printf("s = %T\n", s)

	i3, _ := strconv.Atoi(s1) // 2つ目の戻り値はエラー、_はエラーを無視するための変数
	fmt.Println(i3)

	// エラーハンドリング
	i4, err := strconv.Atoi("A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i4)

	var i7 int = 12345
	s7 := strconv.Itoa(i7)
	fmt.Println(s7)
	fmt.Printf("s7 = %T\n", s7)
}