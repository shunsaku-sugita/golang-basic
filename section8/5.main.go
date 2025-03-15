package main

import "fmt"

// Swtich
// 型スイッチ

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
		case string:
			fmt.Println(v, "!?")
		case int:
			fmt.Println(v + 10000)
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int) // 型アサーション
	fmt.Println(i + 2)
	// fmt.Println(x + 2) // エラー

	// 異なる型でアサーションするとエラーになる
	// エラーを回避するためには、第二引数にbool型を指定する
	// この場合、fには0が入り、isFloat64にはfalseが入る
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)


	// if文で型アサーション
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	} else {
		fmt.Println("I don't know")
	}

	// 型スイッチ
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	// 型スイッチの中で変数を使う
	switch v := x.(type) {
		case bool:
			fmt.Println(v, "bool")
		case int:
			fmt.Println(v, "int")
		case string:
			fmt.Println(v, "string")
		default:
			fmt.Println("I don't know")
	}
}