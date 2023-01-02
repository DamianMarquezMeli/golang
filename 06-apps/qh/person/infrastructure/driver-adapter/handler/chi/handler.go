package chi

import (
	"fmt"
	"net/http"

	"github.com/devpablocristo/golang/06-apps/qh/person/application/port"
)

type PersonChiHandler struct {
	personService port.Service
}

func NewPersonChiHandler(ps port.Service) *PersonChiHandler {
	return &PersonChiHandler{
		personService: ps,
	}
}

func (pch *PersonChiHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	// p := s.gtw.GetPatients()
	// if p == nil || len(p) == 0 {
	// 	p = []*models.Patient{}
	// }
	// web.Success(&p, http.StatusOK).Send(w)

	fmt.Println("hola")
	w.Write([]byte("hola"))
}

func (pch *PersonChiHandler) GetPersonByID(w http.ResponseWriter, r *http.Request) {
	// patientID := chi.URLParam(r, "patientID")
	// id, _ := strconv.ParseInt(patientID, 10, 64)
	// patient, err := s.gtw.GetPatientByID(id)

	// if err != nil {
	// 	web.ErrBadRequest.Send(w)
	// 	return
	// }

	// web.Success(&patient, http.StatusOK).Send(w)
}

func (pch *PersonChiHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	// body := r.Body
	// defer body.Close()
	// var cmd models.CreatePatientCMD
	// err := json.NewDecoder(body).Decode(&cmd)

	// if err != nil {
	// 	web.ErrInvalidJSON.Send(w)
	// 	return
	// }

	// patient, err := s.gtw.CreatePatient(&cmd)

	// if err != nil {
	// 	web.ErrBadRequest.Send(w)
	// 	return
	// }

	// log.Println(patient)
	// web.Success(&patient, http.StatusOK).Send(w)
}
