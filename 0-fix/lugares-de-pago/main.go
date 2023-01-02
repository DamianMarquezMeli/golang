package main

import (
	"fmt"

	"github.com/devpablocristo/lugares_de_pago_ccc/ccccsv"
)

func main() {

	lpv := ccccsv.CargarCsv("viejos")
	lpn := ccccsv.CargarCsv("nuevos")

	fmt.Println("Eliminar:")
	elim := restaDeSlices(lpv, lpn)
	for i := 0; i < len(elim); i++ {
		fmt.Println(elim[i].Direccion, "-", elim[i].Localidad)

	}

	fmt.Println("---------------------------------------")

	fmt.Println("Agregar:")
	agre := restaDeSlices(lpn, lpv)
	for i := 0; i < len(agre); i++ {
		fmt.Println(agre[i].Direccion, "-", agre[i].Localidad)

	}
}

// encuentra los que elementos que solo están en slice1
func restaDeSlices(slice1 ccccsv.LugaresDePago, slice2 ccccsv.LugaresDePago) ccccsv.LugaresDePago {
	var diff ccccsv.LugaresDePago
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1.Comp == s2.Comp {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}
	return diff
}
