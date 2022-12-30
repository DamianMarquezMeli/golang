package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "John", Age: 31}
	p2 := Person{Name: "Mary", Age: 41}
	p3 := Person{Name: "Tom", Age: 66}
	p4 := Person{Name: "Anna", Age: 52}
	p5 := Person{Name: "Paul", Age: 24}
	p6 := Person{Name: "Jane", Age: 33}

	Persons1 := []*Person{&p1, &p2, &p3, &p4, &p1, &p2, &p3, &p5, &p6, &p4, &p5, &p6}
	Persons2 := []*Person{&p4, &p5, &p6, &p1, &p2, &p3, &p1, &p2, &p3, &p5, &p6, &p4, &p5, &p6}
	checkResult := []*Person{&p1, &p2, &p3, &p4, &p5, &p6}

	result := MergePointersSlices(Persons1, Persons2)

	sort.Slice(checkResult, func(i, j int) bool {
		return checkResult[i].Name < checkResult[j].Name
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	fmt.Println("Result:")
	printPointerPersonSlice(result)
	fmt.Println("********************************************************")
	fmt.Println("Expected result:")
	printPointerPersonSlice(checkResult)

}

// Merge slices of pointers to structs without duplicates: map aproach
func MergePointersSlices(s1 []*Person, s2 []*Person) []*Person {

	// create a map to check if a value is already in the slice
	check := make(map[*Person]interface{})

	// add all values from s1 and s2 tp s3
	s1 = append(s1, s2...)

	// create a new slice to store the unique values
	res := make([]*Person, 0)

	// iterate over the slice and check if the value is already in the map
	// if the value is duplicate, it will be overwritten as map key
	for _, per := range s1 {
		check[per] = true
	}

	// iterate over the map and add the unique values to the new slice
	for p, _ := range check {
		res = append(res, p)
	}

	return res
}

func printPointerPersonSlice(persons []*Person) {
	for _, person := range persons {
		println(person.Name, person.Age)
	}
}
