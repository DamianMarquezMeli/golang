package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//investigar como se usa url.path
	fmt.Fprintf(w, "fondin, nana y nano!: %s", r.URL.Path[1:])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//no entiendo pq este no funciona y el otro si
	//obviamente es el mux

	/*

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		        fmt.Fprintf(w, "Hello, you've requesteasdadasd: %s\n", r.URL.Path)
		    })

			http.ListenAndServe(":8080", nil)

	*/

}
