package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	// catNames := []string{"Toribio", "Nana", "Fortunata", "Oki"}
	// err := tpl.Execute(os.Stdout, catNames)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// paste it on the .gohtml file
	// {{range $index, $element := .}}
	// <li>{{$index}}.{{$element}}</h1>
	// {{end}}

	////////////////////

	// countries := map[string]string{
	// 	"Argentina": "America",
	// 	"India":     "Asia",
	// 	"Nigeria":   "Africa",
	// 	"Italia":    "Europe",
	// }
	// err := tpl.Execute(os.Stdout, countries)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// paste it on the .gohtml file
	// 	{{range $key, $val := .}}
	// <li>{{$key}}.{{$val}}</h1>
	// {{end}}

	////////////////////

	// lennon := Person{
	// 	Name: "John",
	// 	Age:  30,
	// }

	// err := tpl.Execute(os.Stdout, lennon)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// paste it on the .gohtml file
	// {{ $n := .Name }}
	// {{ $a := .Age  }}
	// <li>{{$n}} - {{$a}}</h1>

	////////////////////

	lennon := Person{
		Name: "John",
		Age:  30,
	}

	mcartney := Person{
		Name: "Paul",
		Age:  29,
	}

	star := Person{
		Name: "Ringo",
		Age:  25,
	}

	harrison := Person{
		Name: "George",
		Age:  31,
	}

	ps := []Person{lennon, mcartney, star, harrison}

	err := tpl.Execute(os.Stdout, ps)
	if err != nil {
		log.Fatalln(err)
	}

	// paste it on the .gohtml file
	// {{range .}}
	// <li>{{.Name}} - {{.Age}}</h1>
	// {{end}}

}
