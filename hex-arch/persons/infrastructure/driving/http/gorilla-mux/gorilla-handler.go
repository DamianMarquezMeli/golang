package goriadapter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	httpport "github.com/devpablocristo/go-concepts/hex-arch/persons/application/ports/driving"
	domain "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"
)

type GorillaHandler struct {
	service   httpport.PersonService
	MuxRouter *mux.Router
}

type ResponseInfo struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// El handler NO DEBE tener la resposablidad de crear los servicios.
func NewGorillaHandler(s httpport.PersonService) *GorillaHandler {
	r := mux.NewRouter()

	return &GorillaHandler{
		service:   s,
		MuxRouter: r,
	}
}

func (h *GorillaHandler) RunServer(port string) {
	if port == "default" {
		port = ":9000"
	}

	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, h.MuxRouter))
}

func (h *GorillaHandler) SetupRoutes() {
	h.MuxRouter.HandleFunc("/person", h.List).Methods("GET")
	h.MuxRouter.HandleFunc("/person", h.Register).Methods("POST")
}

func (h *GorillaHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPerson domain.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	err = h.service.CreatePerson(newPerson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusCreated,
		Data:   newPerson,
	})
}

func (h *GorillaHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	personsList := h.service.List()

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   personsList,
	})
}
