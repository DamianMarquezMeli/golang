package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-rod/rod"

	domain "github.com/devpablocristo/golang/hex-arch-5/internal/core/domain"
	patientservice "github.com/devpablocristo/golang/hex-arch-5/internal/core/service"
	ginhandler "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/handlers/gin"
	memkvsrepo "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/repositories/kvs"
	gorodservice "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/scrappers/go-rod"
)

const webServerPort string = ":8088"

func main() {

	patPerson := domain.Person{
		UUID:     "1",
		Name:     "Homero",
		Lastname: "Simpson",
		DNI:      12345,
		Gender:   "m",
	}

	docPerson := domain.Person{
		UUID:     "2",
		Name:     "Nick",
		Lastname: "Riviera",
		DNI:      63435,
		Gender:   "m",
	}

	doctor := domain.Doctor{
		Doctor:     docPerson,
		Speciality: "Surgery",
	}

	patient := &domain.Patient{
		Patient:   patPerson,
		Doctor:    doctor,
		Hospital:  "General",
		Diagnosis: "Cancer",
	}

	b, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	patientRepository := memkvsrepo.NewMemKVS()
	patientService := patientservice.NewPatientService(patientRepository)
	goRodService := gorodservice.NewGoRodService(browser)
	patientHandler := ginhandler.NewGinHandler(patientService, goRodService)

	runPatientService(patientHandler)

}
