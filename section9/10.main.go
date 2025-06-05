package main

import "fmt"

// channel
// close

func main() {
	ch1 := make(chan int, 2)

	close(ch1)

	// ch1 <-1

	// fmt.Println(<-ch1)

	i, ok := <-ch1 // 
	fmt.Println(i, ok)
}