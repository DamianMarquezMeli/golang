package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	muxrouter "github.com/devpablocristo/golang/06-apps/bookstore/inventory/infrastructure/router"
)

var httpMuxRouter muxrouter.MuxRouter = *muxrouter.NewMuxRouter()

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/inventory", getBook).Methods("GET")
	router.HandleFunc("/inventory", listBooks).Methods("GET")
	router.HandleFunc("/inventory", addBook).Methods("POST")
	router.HandleFunc("/inventory/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/inventory/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/inventory/{id}", deleteBookByID).Methods("DELETE")
	router.HandleFunc("/inventory/{id}", updateBookByPatch).Methods("PATCH")
	router.HandleFunc("/inventoryISBN/{isbn}", isbn_containes).Methods("GET")

	httpMuxRouter.POST("/inventory/add", inventoryControllers.Add)
	httpMuxRouter.GET("/inventory/all", inventoryControllers.GetAll)
	httpMuxRouter.SERVE(":8888")

	const port string = ":8888"
	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

	return router
}
