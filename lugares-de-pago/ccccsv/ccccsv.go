package ccccsv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/devpablocristo/lugares_de_pago_ccc/ccctxt"
)

//structs

type info struct {
	Direccion string
	Localidad string
	Comp      string
}

type LugaresDePago []info

func CargarCsv(nArch string) LugaresDePago {
	aCsv, err := os.Open(nArch + ".csv")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer aCsv.Close()

	r, err := csv.NewReader(aCsv).ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var lp LugaresDePago
	var aux info
	for _, line := range r {
		aux.Direccion = strings.ToUpper(line[0])
		aux.Localidad = strings.ToUpper(line[1])
		aux.Comp = formatearString(line[0]) + formatearString(line[1])
		lp = append(lp, aux)
	}

	// eliminar primera línea, la cabecera
	lp = lp[1:]

	return lp
}

/*
//maps

type LPago []map[string]string

func CargarCsv(nArch string) LPago {
	aCsv, err := os.Open(nArch + ".csv")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer aCsv.Close()

	r, err := csv.NewReader(aCsv).ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//var lp lugaresDePago
	//var aux info

	var lpMap LPago

	for _, line := range r {
		//aux.Id = i
		//aux.Direccion = formatearString(line[0])
		//aux.Localidad = formatearString(line[1])

		//i++
		//lp = append(lp, aux)

		auxMap := make(map[string]string)
		auxMap[formatearString(line[0])] = formatearString(line[1])
		lpMap = append(lpMap, auxMap)

	}

	// eliminar primera línea, la cabecera
	//lp = lp[1:]
	lpMap = lpMap[1:]

	//return &lp
	return lpMap
}
*/

func formatearString(s string) string {
	sF := strings.ToUpper(s)
	sF = ccctxt.NormalizarTexto(sF)
	sF = ccctxt.AlfaNumerico([]byte(sF))

	return sF
}
