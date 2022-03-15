package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 100

	ch2 := make(chan int, 1)
	ch2 <- 200

	select {
	case n, ok := <-ch1:
		fmt.Println(n, ok)
	case n, ok := <-ch2:
		fmt.Println(n, ok)
	}
}
