package main

import "fmt"

// Panic & Recover

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error!")
	fmt.Println("Start")
}