package main

import "fmt"

func main() {
	c := []int{0, 0, 0, 0, 1, 0}

	fmt.Println(saltarEnNubes(c))
}

func saltarEnNubes(c []int) int {

	//fmt.Println(c)

	saltos := 0
	for i := 0; i < len(c); i++ {
		j := i + 2
		if c[j] == 0 {
			saltos++
			i = j
		} else {
			j := i + 1
			if c[j] == 0 {
				saltos++
				i = j
			}
		}
	}

	return saltos
}
