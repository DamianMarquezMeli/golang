package chiAdapter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	//"github.com/go-chi/chi/v5"

	"github.com/devpablocristo/golang/06-apps/qh/person/application/port"
	"github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type ChiHandler struct {
	personService port.Service
	chiRouter     *chi.Mux
	httpServer    *http.Server
}

func NewChiHandler(ps port.Service, sp string) *ChiHandler {
	cr := chi.NewRouter()

	sv := &http.Server{
		Addr:         ":" + sp,
		Handler:      cr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	//chiRouter.Use("cors")
	//chiRouter.Use(middleware.Logger)

	return &ChiHandler{
		personService: ps,
		chiRouter:     cr,
		httpServer:    sv,
	}
}

func (h *ChiHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	// p := s.gtw.GetPatients()
	// if p == nil || len(p) == 0 {
	// 	p = []*models.Patient{}
	// }
	// web.Success(&p, http.StatusOK).Send(w)

	// code := chi.URLParam(r, "code")
	// redirect, err := useCase.CodeToUrl(code)
	// if err != nil {
	// 	if err == entity.ErrRedirectNotFound {
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 		return
	// 	}
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }
	// http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)

	fmt.Println("hola")
	w.Write([]byte("hola"))
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
	//r.Header().Set("Content-Type", contentType)
	ctx := r.Context()

	body := r.Body
	defer body.Close()
	var newPerson domain.Person
	err := json.NewDecoder(body).Decode(&newPerson)
	if err != nil {
		log.Println("Invalid JSON")
		w.Write([]byte("Invalid JSON"))
		return
	}

	person, err := h.personService.CreatePerson(ctx, &newPerson)
	if err != nil {
		log.Println("Bad Request")
		w.Write([]byte("Bad Request"))
		return
	}

	log.Println(person)
	w.Write([]byte(person.Name))

	w.Header().Set("Content-Type", "application/json")

	//logger.Error("Error decoding JSON", errors.New("test logger"))

	var newLogin domain.Login
	err := json.NewDecoder(r.Body).Decode(&newLogin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(FuryHandlerError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Method:     "furyhandler.Help",
		})
		if err != nil {
			logger.Errorf("Error decoding JSON", err)
			return err
		}
		logger.Errorf("Error parcing body", err)
		return err
	}

	ctx := context.WithValue(r.Context(), userID, newLogin.UserID)
	info, err := hf.autoTestService.CheckHelp(ctx, newLogin)
	if err != nil {
		err := json.NewEncoder(w).Encode(FuryHandlerError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Method:     "autoTestService.CheckHelp",
		})
		logger.Errorf("Error on application layer", err)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(domain.ResponseInfo{
		Status: http.StatusCreated,
		Data:   info,
	})
	return nil

}
