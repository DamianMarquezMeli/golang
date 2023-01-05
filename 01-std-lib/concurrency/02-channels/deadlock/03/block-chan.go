package main

import "fmt"

func main() {
	c := make(chan int, 2)
	d := make(chan int)

	c <- 43
	c <- 34

	go func() {
		d <- 123
	}()

	fmt.Println(<-d)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
