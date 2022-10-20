package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type persona2 struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := persona{
		nombre:   "Homero",
		apellido: "Simpson",
		edad:     39,
	}

	p2 := persona{
		nombre:   "Marge",
		apellido: "Simpson",
		edad:     39,
	}

	p3 := persona2{
		Nombre:   "Homero",
		Apellido: "Simpson",
		Edad:     39,
	}

	p4 := persona2{
		Nombre:   "Marge",
		Apellido: "Simpson",
		Edad:     39,
	}

	gente := []persona{
		p1,
		p2,
	}

	fmt.Println(gente)

	// funciona pero esta vacío
	// pq los campos de persona esta en minúscula
	bs, err := json.Marshal(gente)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	gente2 := []persona2{
		p3,
		p4,
	}

	fmt.Println(gente2)

	// funciona correctamente
	// porque los campos de persona2 empiezan con mayuscula

	// lo que hace esto, es transformar el struct en un json
	// para por poder enviarlo

	bs, err = json.Marshal(gente2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
