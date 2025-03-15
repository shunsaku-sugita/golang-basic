package main

import "fmt"

// Switch


func main() {

	/*
	n := 0
	switch n {
	case 0, 1, 2:
		fmt.Println("0, 1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		println("I don't know")
	}
	*/

	// 短縮文
	/*
	switch n :=2; n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			println("I don't know")
	}
			*/

	n := 6
	switch {
		case n > 0 && n < 4:
			fmt.Println("0 < n < 4")
		case n > 3 && n < 7:
			fmt.Println("3 < n < 7")
		default:
			fmt.Println("I don't know")
	}

}