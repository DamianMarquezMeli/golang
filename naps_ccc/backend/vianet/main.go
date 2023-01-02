package main

import (
	"fmt"
	"log"
	"net/http"

	"vianet/controllers"

	"github.com/gorilla/mux"
)

func main() {

	controller := controllers.Controller{}

	r := mux.NewRouter()

	r.HandleFunc("/", controller.Index())
	// /vianet/descripcion
	//r.HandleFunc("/vianet/{id}", controller.GetNaps()).Methods("GET")
	r.HandleFunc("/vianet", controller.GetNap()).Methods("GET")

	puerto := "8003"
	fmt.Println("El servidor corre en el puerto:", puerto)
	err := http.ListenAndServe(":"+puerto, r)
	logErrores(err)

}

func logErrores(err error) {
	if err != nil {
		log.Fatal("ERROR!!!", err)
	}
}
