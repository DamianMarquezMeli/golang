package chihandler

import (
	"encoding/json"
	"log"
	"net/http"

	cdomain "github.com/devpablocristo/golang/06-apps/qh/internal/commons/domain"
	port "github.com/devpablocristo/golang/06-apps/qh/person/application/port"
	domain "github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type ChiHandler struct {
	personService port.Service
}

func NewChiHandler(ps port.Service) *ChiHandler {
	return &ChiHandler{
		personService: ps,
	}
}

func (h *ChiHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var errReq cdomain.APIError
	errReq.Method = "chiAdapter.GetPerson"

	ctx := r.Context()
	persons, err := h.personService.GetPersons(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errReq = cdomain.ErrInternalServer
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			cdomain.ErrInternalServer,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		cdomain.ResponseAPI{
			Success: true,
			Status:  http.StatusCreated,
			Result:  persons,
		},
	)

	// fmt.Println("hola")
	// w.Write([]byte("hola"))
}

func (h *ChiHandler) GetPersonByID(w http.ResponseWriter, r *http.Request) {
	// patientID := chi.URLParam(r, "patientID")
	// id, _ := strconv.ParseInt(patientID, 10, 64)
	// patient, err := s.gtw.GetPatientByID(id)

	// if err != nil {
	// 	web.ErrBadRequest.Send(w)
	// 	return
	// }

	// web.Success(&patient, http.StatusOK).Send(w)
}

func (h *ChiHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close()

	var errReq cdomain.APIError
	errReq.Method = "chiAdapter.CreatePerson"

	var newPerson domain.Person
	err := json.NewDecoder(body).Decode(&newPerson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errReq = cdomain.ErrInvalidJSON
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			errReq,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	ctx := r.Context()
	person, err := h.personService.CreatePerson(ctx, &newPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errReq = cdomain.ErrInternalServer
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			cdomain.ErrInternalServer,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		cdomain.ResponseAPI{
			Success: true,
			Status:  http.StatusCreated,
			Result:  person,
		},
	)
}
