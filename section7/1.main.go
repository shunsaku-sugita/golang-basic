package main

import "fmt"

// ========================================
// 関数
func plus(x int, y int) int {
	return x + y
}

func div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 戻り値の変数名を指定すると、returnの後に変数名を指定することで、returnの後に何も書かなくても返り値を返すことができる
func double(price int) (result int) {
	result = price * 2
	return
}

func noReturn() {
	fmt.Println("No return")
}

func main() {
	i := plus(1, 2)
	fmt.Println(i)

	i2, _ := div(9, 3)
	fmt.Println(i2)

	i4 := double(1000)
	fmt.Println(i4)

	noReturn()
}