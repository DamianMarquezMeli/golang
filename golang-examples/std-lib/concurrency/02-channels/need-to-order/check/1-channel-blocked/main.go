package main

import "fmt"

func main() {
	portal := make(chan string)

	portal <- "Ironman"

	fmt.Println(<-portal)
}
