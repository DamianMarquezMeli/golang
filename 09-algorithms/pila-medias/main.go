package main

import "fmt"

func main() {

	// ar es una pila de medias, cada numero es un media, encotrar cuantos pares hay
	ar := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(paresEnArray(ar))
}

func paresEnArray(arr []int) int {
	//Create a   dictionary of values for each element
	m := make(map[int]int)

	pares := 0
	for _, num := range arr {
		m[num] += 1

		if m[num]%2 == 0 {
			pares++
			m[num] = 0
		}
	}
	//fmt.Println(m)

	return pares
}
